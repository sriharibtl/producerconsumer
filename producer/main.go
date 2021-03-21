package main

import (
	"httpserver"
)

const (
	ip = "0.0.0.0:9990"
)

var counter int

func main() {
	panic(httpserver.StartHttpServer(ip))
}
