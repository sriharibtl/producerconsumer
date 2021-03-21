package main

import (
	"context"
	"httpserver"
	"kafka"
)

const (
	ip = "0.0.0.0:9991"
)

var counter = 10

func main() {
	go kafka.Consume(context.Background())
	panic(httpserver.StartHttpServer(ip))
}
