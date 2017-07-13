package main

import (
	"fmt"
	model "go-rest-server/model"
	"net/http"
	"strconv"
)

type RestServer struct {
	Port     int
	Postfix  string
	Handlers []*model.HTTPHandler
}

func CreateRestServer(port int) *RestServer {
	return &RestServer{port, "", nil}
}

func (rs *RestServer) Handle(handlers ...*model.HTTPHandler) *RestServer {
	rs.Handlers = handlers
	return rs
}

func (s *RestServer) Start(contextPath string) {

	portString := strconv.Itoa(s.Port)
	if s.Port <= 0 || s.Port >= 65536 {
		panic("Port cannot be " + portString)
	}

	portString = ":" + portString

	for _, handler := range s.Handlers {
		http.HandleFunc(contextPath+handler.URI, handler.Handler)
	}
	fmt.Println(portString)
	http.ListenAndServe(portString, nil)
}
