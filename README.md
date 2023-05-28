# TCP Chat with Golang

## Description

This is a simple TCP chat server and client written in Golang. The server is able to handle multiple clients at once and broadcast messages to all connected clients. The client is able to send messages to the server and receive messages from the server.

## Usage

### Server

To run the server, run the following command:

```
go run serverMain/server.go
```

The server and client are currently set to run on localhost:3090. This can be changed with the flags `--host=<your_host> --port=<your_port>`.
This way you can run many server on a remote machine and connect to it with the client.

### Client

To run the client, run the following command:

```
go run chatGo/chat.go
```

The server and client are currently set to run on localhost:3090. This can be changed with the flags `--host=<your_host> --port=<your_port>` to connect with a especific server.

## Note

This repository include a build executable for Linux of the files server and chat.
This is a simple project to learn more about Golang. It is not intended to be used in production.
