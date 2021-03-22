package kafka

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

var (
	data      = "kafka!"
	KafkaHost string
	Topic     string
	w         *kafka.Writer
)

func Produce(ctx context.Context) {
	// intialize the writer with the broker addresses, and the topic
	if w == nil {
		w = kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{KafkaHost},
			Topic:   Topic,
		})
	}

	err := w.WriteMessages(ctx, kafka.Message{
		// Key: []byte(strconv.Itoa(1)),
		// create an arbitrary message payload for the value
		Value: []byte("this is message " + data),
	})
	if err != nil {
		log.Println("could not write message " + err.Error())
		return
	}
	log.Println("Meassage written")

}
