package model

import (
	"encoding/json"
	"net/http"
)

// ErrorJSON use like error response in JSON
type ErrorJSON struct {
	Error string `json:"error"`
}

// HTTPContext groups incoming values
type HTTPContext struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// GetBody return scpecific JSON body
func (c *HTTPContext) GetBody(bodyHolder interface{}) (interface{}, error) {
	decoder := json.NewDecoder(c.Request.Body)
	return bodyHolder, decoder.Decode(bodyHolder)
}

// WriteBody write JSON body to response
func (c *HTTPContext) WriteBody(bodyHolder interface{}) error {
	c.Response.Header().Add("Content-Type", "application/json;charset=UTF-8")
	encoder := json.NewEncoder(c.Response)
	return encoder.Encode(bodyHolder)
}

// SetHeaders set response headers
func (c *HTTPContext) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		c.Response.Header().Add(key, value)
	}
}

// WriteError write JSON error to response
func (c *HTTPContext) WriteError(message string, statusCode HTTPStatusCode) {
	c.Response.WriteHeader(int(statusCode))
	c.WriteBody(ErrorJSON{message})
}
