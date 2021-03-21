package httpserver

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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
		Route{Name: "musicGet", Path: "/counter", Method: "GET", Handler: c.getHandler},
	}
}

var GetCounter func() int

func (c *controller) getHandler(resp http.ResponseWriter, req *http.Request) {
	log.Println("Received request")
	resp.Write([]byte(strconv.Itoa(GetCounter())))
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
