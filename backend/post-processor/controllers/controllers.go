package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	kafka_messages "my-posting-site/backend/common/kafka"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	topic           = "post-processing"
	topic_processed = "post-processed"
	broker1Address  = "localhost:9092"
	broker2Address  = "localhost:9093"
)

func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topic,
		GroupID: "post-processor-group",
	})

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topic_processed,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		post := &kafka_messages.Post{}

		fmt.Println("received: ", string(msg.Value))
		json.Unmarshal(msg.Value, &post)

		post = ProcessPost(post)
		marshalled, _ := json.Marshal(post)

		err = writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(""),
			Value: marshalled,
		})

		if err != nil {
			panic("could not write message " + err.Error())
		}

	}
}

func ProcessPost(post *kafka_messages.Post) *kafka_messages.Post {
	for _, elem := range post.Elements {
		text := strings.Trim(*elem.Text, " ")
		title := strings.Trim(*elem.Title, " ")
		elem.Text = &text
		elem.Title = &title
	}

	return post
}
