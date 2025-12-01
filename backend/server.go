package backend

import "github.com/bitwormhole/passwordbox/core/pemapi"

type Server struct {
	foo string

	codec pemapi.CODEC
}

// ExecuteMessage implements pemapi.API.
func (inst *Server) ExecuteMessage(request *pemapi.Request, response *pemapi.Response) error {
	panic("unimplemented")
}

// Execute implements pemapi.API.
func (inst *Server) Execute(request *pemapi.Request, response *pemapi.Response) error {
	panic("unimplemented")
}

// ExecuteContext implements pemapi.API.
func (inst *Server) ExecuteContext(c *pemapi.Context) error {
	panic("unimplemented")
}

// ExecuteString implements pemapi.API.
func (inst *Server) ExecuteString(request string) (response string) {
	panic("unimplemented")
}

// ExecuteText implements pemapi.API.
func (inst *Server) ExecuteText(request pemapi.RequestText) (response pemapi.ResponseText) {
	panic("unimplemented")
}

func NewServer() *Server {
	return new(Server)
}

func (inst *Server) _impl() pemapi.API {
	return inst
}
