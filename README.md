# JOBSITY CHAT SERVER
This repository houses a chat server that employs websockets and RabbitMQ for message communication. 
All data is persisted in a PostgreSQL database.

### Dependencies
This repo is powered by RabbitMQ and PostgreSQL.
Both are managed using docker-compose.

### Running locally

To run the server execute: `make run`
All configurations for RabbitMQ and PostgreSQL are presented in the `Makefile` and in the `docker-compose.yml` file.

### API 

These are the endpoints provided by the server:
```
bash 
| Endpoint                              | HTTP Methods | Params                         | Description                                      |
| :-------------------------------------| :----------- | :----------------------------- | :----------------------------------------------- |
| /v1/auth/signup                       | POST         | `username` `email` `password`  | Creates a new user and returns jwt session token |
| /v1/auth/login                        | POST         | `email` `password`             | Logs in a user and returns the jwt session token |
| /v1/chat/rooms                        | POST         | `name`                         | Creates a new chat room with the name provided   |
| /v1/chat/rooms                        | GET          |                                | returns a list of chat rooms                     |
| /v1/api/chat/rooms/:id/messages       | GET          |                                | Returns the latest 50 messages                   |
| /v1/ws                                | GET          | token                          | websocket connection url                         |
```

For the websocket the expected message entry like this:
``````
{
    "type":1,
    "chatmessage":"/stock=APPL.US",
    "chatuser":"wesmota@gmail.com",
    "chatroomId":1
}
``````
### Migrations 

Database migrations are managed by [Goose](https://github.com/pressly/goose). Migrations live in the root directory.

### Running your migration locally.

`make db-reset` will create a database from the current schema and run any new migrations. See the 
Makefile for details. 

`make db-migrate`, which is called by `make db-reset` dumps the DB schema.sql file after all migrations
have been run. For any new migrations you MUST submit the updated schema.sql file. We use that file to 
quickly build databases locally.

There are a few more commands in the Makefile that you can use when working with the local DB to build migrations. Refer to the Goose documentation for seeing how to use that command directly.

## Postman
Postman + terminal logs were used to test the websocket connection and the server. The file used was exported and it's inside the `postman`` folder.
