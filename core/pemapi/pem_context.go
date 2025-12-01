package pemapi

type Context struct {
	Request *Request

	Response *Response

	Key HandlerKey

	Handler Handler

	Filters FilterChain
}
