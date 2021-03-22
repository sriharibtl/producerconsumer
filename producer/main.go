package main

import (
	"errors"
	"httpserver"
	"kafka"
	"os"
)

const (
	ip          = "0.0.0.0:9990"
	KAFKA_HOST  = "KAFKA_HOST"
	KAFKA_TOPIC = "KAFKA_TOPIC"
)

func LoadConfigFromEnv() error {

	kafkahost := os.Getenv(KAFKA_HOST)
	if kafkahost == "" {
		return errors.New("KAFKA HOST IP Environment variable not set")
	}
	kafka.KafkaHost = kafkahost

	kafkatopic := os.Getenv(KAFKA_TOPIC)
	if kafkatopic == "" {
		return errors.New("KAFKA TOPIC Environment variable not set")
	}
	kafka.Topic = kafkatopic

	return nil
}

func main() {
	err := LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	panic(httpserver.StartHttpServer(ip))
}
