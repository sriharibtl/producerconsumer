package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"

	kafka "kafka"

	"github.com/gorilla/mux"
)

type Route struct {
	Method, Name, Path string
	Handler            http.HandlerFunc
}

type Router interface {
	Routes() []Route
}

func NewRouter(routes ...Router) *mux.Router {
	router := mux.NewRouter()
	for _, api := range routes {
		for _, route := range api.Routes() {
			router.Methods(route.Method).Name(route.Name).Path(route.Path).Handler(route.Handler)
		}
	}
	return router
}

type controller struct {
}

func (c *controller) Routes() []Route {
	return []Route{
		Route{Name: "musicGet", Path: "/incrementcounter", Method: "GET", Handler: c.getHandler},
	}
}

func (c *controller) getHandler(resp http.ResponseWriter, req *http.Request) {
	log.Println("Received request")
	if !kafka.ProducerStarted {
		go kafka.Produce(context.Background())
		kafka.ProducerStarted = true
	}
	resp.Write([]byte("Producer started"))
}

//StartHttpServer -Start Http server using the router
func StartHttpServer(ip string) error {
	if ip == "" {
		return errors.New("Empty IP")
	}
	controllerObj := &controller{}
	router := NewRouter(controllerObj)
	server := http.Server{Addr: ip, Handler: router}
	log.Println("Starting server on ip:", ip)
	return server.ListenAndServe()
}
