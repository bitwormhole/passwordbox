package pemapi

type StringHandlerFunc func(request string) (response string)

type TextHandlerFunc func(request RequestText) (response ResponseText)

type MessageHandlerFunc func(request *Request, response *Response) error

type ContextHandlerFunc func(c *Context) error

type API struct {
	StringHandler  StringHandlerFunc
	TextHandler    TextHandlerFunc
	MessageHandler MessageHandlerFunc
	ContextHandler ContextHandlerFunc
}

type HandlerFunc = ContextHandlerFunc
