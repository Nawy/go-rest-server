package model

// HTTPMethodType is HTTP methods type for controlling params by type
type HTTPMethodType string

// HTTPStatusCode is HTTP status code type for controlling params by type
type HTTPStatusCode int

const (
	// GET method requests a representation of the specified resource. Requests using GET should only retrieve data.
	GET = HTTPMethodType("GET")
	// HEAD method asks for a response identical to that of a GET request, but without the response body.
	HEAD = HTTPMethodType("HEAD")
	// POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server
	POST = HTTPMethodType("POST")
	// PUT method replaces all current representations of the target resource with the request payload.
	PUT = HTTPMethodType("PUT")
	// DELETE method deletes the specified resource.
	DELETE = HTTPMethodType("DELETE")
	// CONNECT method establishes a tunnel to the server identified by the target resource.
	CONNECT = HTTPMethodType("CONNECT")
	// OPTIONS method is used to describe the communication options for the target resource.
	OPTIONS = HTTPMethodType("OPTIONS")
	// TRACE method performs a message loop-back test along the path to the target resource.
	TRACE = HTTPMethodType("TRACE")
	// PATCH method is used to apply partial modifications to a resource.
	PATCH = HTTPMethodType("PATCH")
)

const (
	HTTP_OK          = HTTPStatusCode(200)
	HTTP_BAD_REQUEST = HTTPStatusCode(400)
	UNATHORIZED      = HTTPStatusCode(401)
	HTTP_NOT_FOUND   = HTTPStatusCode(404)

	HTTP_INTERNAL_SERVER_ERROR = HTTPStatusCode(500)
	HTTP_NOT_IMPLEMENTED       = HTTPStatusCode(501)
)
