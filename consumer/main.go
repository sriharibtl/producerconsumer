package main

import (
	"cbapi"
	"context"
	"errors"
	"httpserver"
	"kafka"
	"os"
)

const (
	ip          = "0.0.0.0:9991"
	DB_IP       = "DB_IP"
	KAFKA_HOST  = "KAFKA_HOST"
	KAFKA_TOPIC = "KAFKA_TOPIC"
)

func LoadConfigFromEnv() error {

	couchbaseip := os.Getenv(DB_IP)
	if couchbaseip == "" {
		return errors.New("DB IP Environment variable not set")
	}
	cbapi.DB_IP = couchbaseip

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
	err = cbapi.InitDB()
	if err != nil {
		panic(err)
	}
	go kafka.Consume(context.Background())
	panic(httpserver.StartHttpServer(ip))
}
