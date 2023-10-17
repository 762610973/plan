package main

import (
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
	queue, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil {
		slog.Error("failed to declare a queue", "err:", err.Error())
		return
	}
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		slog.Error("failed to consumer from queue", err.Error())
		return
	}
	for {
		v := <-msgs
		slog.Info("receive message: ", slog.String("value", string(v.Body)))
	}
}
