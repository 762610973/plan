package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"math/rand"
	"strconv"
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
	// 声明交换机
	//ch.ExchangeDeclare()
	// 声明队列
	//ch.QueueDeclare()
	// 绑定
	//ch.QueueBind()
	body := []byte(strconv.FormatInt(rand.Int63(), 10))
	var bindingKey string
	err = ch.PublishWithContext(
		context.Background(),
		"test", bindingKey, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
	if err != nil {
		slog.Error("publish failed", err.Error())
		return
	}
	slog.Info("send success", "the message is", string(body))
}
