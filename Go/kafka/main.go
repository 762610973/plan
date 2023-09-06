package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"log/slog"
	"time"
)

const (
	_address = "192.168.10.221:9092"
	//_address  = "192.168.31.22:9092"
	_username = "admin"
	_password = "admin"
	_topic    = "first"
)

func main() {
	mechanism := plain.Mechanism{
		Username: _username,
		Password: _password,
	}
	exitCh := make(chan struct{})
	go func() {
		writer := kafka.Writer{
			Addr:  kafka.TCP(_address),
			Topic: _topic,
			Transport: &kafka.Transport{
				SASL:        mechanism,
				DialTimeout: 3 * time.Second,
			},
			AllowAutoTopicCreation: true,
		}
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte("write to kafka"),
		})
		if err != nil {
			slog.Error("write failed", err.Error())
			return
		}
		slog.Info("write to kafka success")
		_ = writer.Close()
	}()
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			StartOffset: kafka.LastOffset,
			Brokers:     []string{_address},
			Topic:       _topic,
			Dialer: &kafka.Dialer{
				DualStack:     true,
				SASLMechanism: mechanism,
			},
			MaxBytes: 10e6,
		})

		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			slog.Error("read failed", err.Error())
			return
		}
		if message.Value == nil {
			return
		}
		slog.Info(string(message.Value))
		_ = reader.Close()
		exitCh <- struct{}{}
	}()
	<-exitCh
}
