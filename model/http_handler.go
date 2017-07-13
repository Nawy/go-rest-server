package model

import "net/http"

type HTTPHandlerFunc func(http.ResponseWriter, *http.Request)

type HTTPHandler struct {
	Method  HTTPMethodType
	URI     string
	Handler HTTPHandlerFunc
}

func CreateHTTPHandler(method HTTPMethodType, uri string) *HTTPHandler {
	h := &HTTPHandler{method, uri, nil}
	h.Handler = func(w http.ResponseWriter, r *http.Request) {
		if !h.isCorrectMethod(r) {
			panic("Wrong http method, you should use " + string(h.Method))
		}
	}
	return h
}

func (h *HTTPHandler) HandleOp(handler HTTPHandlerFunc) *HTTPHandler {
	beforeHandler := h.Handler
	h.Handler = func(w http.ResponseWriter, r *http.Request) {
		beforeHandler(w, r)
		handler(w, r)
	}
	return h
}

func (h *HTTPHandler) HandleOps(handlers ...HTTPHandlerFunc) *HTTPHandler {
	beforeHandler := h.Handler
	h.Handler = func(w http.ResponseWriter, r *http.Request) {
		beforeHandler(w, r)
		for _, handler := range handlers {
			handler(w, r)
		}
	}
	return h
}

func (h *HTTPHandler) isCorrectMethod(r *http.Request) bool {
	if r.Method == string(h.Method) {
		return true
	}

	return false
}
