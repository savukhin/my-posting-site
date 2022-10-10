package post_controllers

import (
	"context"
	"encoding/json"
	"fmt"
	kafka_messages "my-posting-site/backend/common/kafka"

	"github.com/segmentio/kafka-go"
)

const (
	topic           = "post-processing"
	topic_processed = "post-processed"
	broker1Address  = "localhost:9091"
	broker2Address  = "localhost:9092"
	broker3Address  = "localhost:9093"
)

func Produce(post *kafka_messages.Post) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
	})

	marshalled, _ := json.Marshal(post)

	return writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(""),
		Value: marshalled,
	})
}

func Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic_processed,
		GroupID: "post-processed-group",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}

		post := &kafka_messages.Post{}

		fmt.Println("received: ", string(msg.Value))
		json.Unmarshal(msg.Value, &post)

		err = CreatePost(post)
		if err != nil {
			fmt.Println(err)
		}
	}
}
