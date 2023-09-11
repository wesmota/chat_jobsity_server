package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/websocket"
)

type Broker struct {
	ReceiverQueue  amqp.Queue
	PublisherQueue amqp.Queue
	Channel        *amqp.Channel
}

// Setup creates(or connects if not existing) the reciever and publisher queues
func (b *Broker) Setup(ch *amqp.Channel) {
	//based on https://www.rabbitmq.com/tutorials/tutorial-one-go.html

	receiverQueue := "JOBSITY_PUBLISHER"
	publisherQueue := "JOBSITY_RECEIVER"

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
		log.Err(err).Msg("ReadMessages Error occured")
		return
	}

	receivedMsgs := make(chan models.ChatMessage)
	go toMsgResponse(msgs, receivedMsgs)
	go processAndPublish(receivedMsgs, b, hub)
	// keep it running
	select {}
}

func toMsgResponse(entries <-chan amqp.Delivery, receivedMessages chan models.ChatMessage) {
	log.Info().Msg("Converting messages")
	var msgR models.ChatMessage
	for d := range entries {
		log.Info().Msgf("Received a message: %s", d.Body)
		err := json.Unmarshal([]byte(d.Body), &msgR)
		if err != nil {
			log.Err(err).Msg("Error on unmarshalling message")
			continue
		}
		log.Info().Msgf("Received a message: %+v", msgR)
		receivedMessages <- msgR
	}
}

func processAndPublish(msgs <-chan models.ChatMessage, b *Broker, hub *websocket.Hub) {
	log.Info().Msg("Processing messages")
	for m := range msgs {
		log.Info().Msgf("Processing message %s for room: %d", m.ChatMessage, m.ChatRoomId)
		// send message to publisher queue
		chatMsg := models.ChatMessage{
			Type:        1,
			ChatMessage: m.ChatMessage,
			ChatUser:    m.ChatUser,
			ChatRoomId:  m.ChatRoomId,
		}
		hub.Broadcast <- chatMsg
		log.Info().Msg("Message sent to publisher queue")
	}
}
