package pemapi

type API interface {

	// 这是 'ExecuteMessage' 的别名
	Execute(request *Request, response *Response) error

	ExecuteString(request string) (response string)

	ExecuteText(request RequestText) (response ResponseText)

	ExecuteMessage(request *Request, response *Response) error

	ExecuteContext(c *Context) error

	// StringHandler  StringHandlerFunc
	// TextHandler    TextHandlerFunc
	// MessageHandler MessageHandlerFunc
	// ContextHandler ContextHandlerFunc

}

// type HandlerFunc = ContextHandlerFunc
