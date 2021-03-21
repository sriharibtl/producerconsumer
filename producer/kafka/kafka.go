package kafka

import (
	"context"
	"log"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic          = "parking"
	broker1Address = "172.16.8.115:9092"
)

var data = "kafka!"

var w *kafka.Writer

func Produce(ctx context.Context) {
	// initialize a counter
	// i := 0

	// intialize the writer with the broker addresses, and the topic
	if w == nil {
		w = kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{broker1Address},
			Topic:   topic,
		})
	}

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(1)),
		// create an arbitrary message payload for the value
		Value: []byte("this is message " + data),
	})
	if err != nil {
		log.Println("could not write message " + err.Error())
		return
	}
	log.Println("Meassage written")

	// log a confirmation once the message is written
	// fmt.Println("writes:", i)
	// i++
	// sleep for a second
	time.Sleep(time.Second)
}
