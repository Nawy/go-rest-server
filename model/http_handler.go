package model

import "net/http"

// ResponseEntity groups response
type ResponseEntity struct {
	Body    interface{}
	Headers map[string]string
}

// HTTPFilterFunc type for filter function
type HTTPFilterFunc func(*HTTPContext)

// HTTPHandlerFunc type for handler function
type HTTPHandlerFunc func(*HTTPContext) *ResponseEntity

// HTTPHandler handler object
type HTTPHandler struct {
	Method         HTTPMethodType
	URI            string
	IncomeFilters  []HTTPFilterFunc
	OutcomeFilters []HTTPFilterFunc
	Handler        HTTPHandlerFunc
	Dispatcher     func(http.ResponseWriter, *http.Request)
}

// MakeRequestHandler make handlers chain
func MakeRequestHandler(method HTTPMethodType, uri string) *HTTPHandler {
	h := &HTTPHandler{method, uri, nil, nil, nil, nil}
	h.Dispatcher = func(response http.ResponseWriter, request *http.Request) {

		context := &HTTPContext{response, request}

		if !h.isCorrectMethod(request) {
			context.WriteError("Wrong http method, you should use "+string(h.Method), HTTP_BAD_REQUEST)
			return
		}

		for _, incomeFilter := range h.IncomeFilters {
			incomeFilter(context)
		}

		responseEntity := h.Handler(context)

		if responseEntity != nil && responseEntity.Body != nil {
			context.WriteBody(&responseEntity.Body)
		}

		for _, outcomeFilter := range h.OutcomeFilters {
			outcomeFilter(context)
		}
	}
	return h
}

// SetHandler sets handler
func (h *HTTPHandler) SetHandler(handler HTTPHandlerFunc) *HTTPHandler {
	h.Handler = handler
	return h
}

// SetIncome set income filters, remember order is importent
func (h *HTTPHandler) SetIncome(filter ...HTTPFilterFunc) *HTTPHandler {
	h.IncomeFilters = filter
	return h
}

// SetOutcome set income filters, remember order is importent
func (h *HTTPHandler) SetOutcome(filter ...HTTPFilterFunc) *HTTPHandler {
	h.OutcomeFilters = filter
	return h
}

// isCorrectMethod checks HTTP method
func (h *HTTPHandler) isCorrectMethod(r *http.Request) bool {
	if r.Method == string(h.Method) {
		return true
	}
	return false
}

// FullResponse create response with body and headers
func FullResponse(body interface{}, headers map[string]string) *ResponseEntity {
	return &ResponseEntity{Body: body, Headers: headers}
}

// BodyResponse create body response
func BodyResponse(body interface{}) *ResponseEntity {
	return &ResponseEntity{Body: body, Headers: nil}
}

// HeaderResponse create header response
func HeaderResponse(headers map[string]string) *ResponseEntity {
	return &ResponseEntity{Body: nil, Headers: headers}
}
