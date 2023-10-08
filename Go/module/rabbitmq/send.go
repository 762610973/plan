package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.31.22:5672")
	if err != nil {
		slog.Error("failed to connect to rabbitmq", err.Error())
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		slog.Error("failed to open a channel", err.Error())
		return
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		slog.Error("failed to declare a queue", err.Error())
		return
	}
	body := "HELLO WORLD!"
	err = ch.PublishWithContext(
		context.Background(),
		"", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		slog.Error("publish failed", err.Error())
		return
	}
	slog.Info("send success", "the message is", "hello world!")
}
