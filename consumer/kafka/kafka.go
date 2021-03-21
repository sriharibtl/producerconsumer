package kafka

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic          = "parking"
	broker1Address = "172.16.8.115:9092"
)

var data = "kafka!"

var r *kafka.Reader

var counter = 0

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	if r == nil {
		r = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker1Address},
			Topic:   topic,
			GroupID: "my-group",
		})
	}

	// the `ReadMessage` method blocks until we receive the next event
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		counter++
	}

}

func GetCounter() int {
	return counter
}
