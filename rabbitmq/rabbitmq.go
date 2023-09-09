package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/websocket"
)

type Broker struct {
	ReceiverQueue  amqp.Queue
	PublisherQueue amqp.Queue
	Channel        *amqp.Channel
}

type MessageResponse struct {
	RoomId  uint   `json:"RoomId"`
	Message string `json:"Message"`
}

// Setup creates(or connects if not existing) the reciever and publisher queues
func (b *Broker) Setup(ch *amqp.Channel) {
	//based on https://www.rabbitmq.com/tutorials/tutorial-one-go.html

	receiverQueue := "JOBSITY_RECEIVER"
	publisherQueue := "JOBSITY_PUBLISHER"

	qR, err := ch.QueueDeclare(
		receiverQueue, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Info().Msg("Receiver queue already exists")
		return
	}

	qP, err := ch.QueueDeclare(
		publisherQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Info().Msg("Publisher queue already exists")
		return
	}

	b.ReceiverQueue = qR
	b.PublisherQueue = qP
	b.Channel = ch
}

// Publish sends messages to receiver queue
func (b *Broker) Publish(message chan []byte) {
	for body := range message {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err := b.Channel.PublishWithContext(ctx,
			"",                    // exchange
			b.PublisherQueue.Name, // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		cancel()
		if err != nil {
			log.Err(err).Msg("Failed to publish a message")
			continue
		}
		log.Info().Msg("Published a message")
	}
}

// Read reads messages from receiver queue
func (b *Broker) Read(hub *websocket.Hub) {
	msgs, err := b.Channel.Consume(
		b.ReceiverQueue.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Printf("ReadMessages Error occured %s\n", err)
		return
	}

	receivedMsgs := make(chan MessageResponse)
	go func() {
		var msg MessageResponse
		for d := range msgs {
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Err(err).Msg("Failed to unmarshal message")
				continue
			}
			log.Printf("Received a message: %s", d.Body)
			receivedMsgs <- msg
		}
	}()
	// keep it running
	select {}
}
