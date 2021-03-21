package main

import (
	"httpserver"
)

const (
	ip = "0.0.0.0:9991"
)

var counter = 10

func main() {

	httpserver.GetCounter = func() int {
		return counter
	}

	panic(httpserver.StartHttpServer(ip))
}
