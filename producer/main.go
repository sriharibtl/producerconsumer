package main

import (
	"httpserver"
)

const (
	ip = "0.0.0.0:9990"
)

var counter int

func main() {
	httpserver.IncrementHandler = func() {
		counter++
	}

	httpserver.GetCounter = func() int {
		return counter
	}

	panic(httpserver.StartHttpServer(ip))
}
