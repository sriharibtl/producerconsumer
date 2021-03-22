package main

import (
	"cbapi"
	"context"
	"httpserver"
	"kafka"
)

const (
	ip = "0.0.0.0:9991"
)

var counter = 10

func main() {
	err := cbapi.InitDB()
	if err != nil {
		panic(err)
	}
	go kafka.Consume(context.Background())
	panic(httpserver.StartHttpServer(ip))
}
