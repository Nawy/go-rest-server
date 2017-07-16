package rest

import (
	"fmt"
	model "go-rest-server/model"
	"net/http"
	"strconv"
	"strings"
)

type RestServer struct {
	Port       int
	ContextURI string
	Handlers   []*model.HTTPHandler
}

func CreateRestServer(port int) *RestServer {
	return &RestServer{port, "", nil}
}

func (rs *RestServer) SetHandlers(handlers ...*model.HTTPHandler) *RestServer {
	rs.Handlers = handlers
	return rs
}

func (s *RestServer) SetContextURI(URI string) *RestServer {
	s.ContextURI = getCorrectURI(URI)
	return s
}

func (s *RestServer) Start() {

	portString := strconv.Itoa(s.Port)
	if s.Port <= 0 || s.Port >= 65536 {
		panic("Port cannot be " + portString)
	}

	portString = ":" + portString

	for _, handler := range s.Handlers {
		http.HandleFunc(s.ContextURI+handler.URI, handler.Dispatcher)
	}
	fmt.Println(portString)
	http.ListenAndServe(portString, nil)
}

func getCorrectURI(URI string) string {
	if URI == "" {
		return URI
	}

	resultURI := strings.TrimSpace(URI)

	if !strings.HasPrefix(resultURI, "/") {
		resultURI = "/" + resultURI
	}

	if strings.HasSuffix(resultURI, "/") {
		resultURI = strings.TrimSuffix(resultURI, "/")
	}

	return resultURI
}
