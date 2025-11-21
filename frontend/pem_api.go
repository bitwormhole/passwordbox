package frontend

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/backend"
	"github.com/bitwormhole/passwordbox/core/pemapi"
)

func GetAPI() *pemapi.API {
	api := new(pemapi.API)
	api.StringHandler = innerSendString
	api.TextHandler = innerSendText
	api.MessageHandler = innerSendMessage
	api.ContextHandler = innerSendContext
	return api
}

func innerSendString(req string) (resp string) {
	return backend.ExecuteString(req)
}

func innerSendText(req pemapi.RequestText) (resp pemapi.ResponseText) {
	str1 := req.String()
	str2 := innerSendString(str1)
	return pemapi.ResponseText(str2)
}

func innerSendMessage(req *pemapi.Request, resp *pemapi.Response) error {

	if req == nil || resp == nil {
		return fmt.Errorf("param (req|resp) is nil")
	}

	t1 := req.Text()
	t2 := innerSendText(t1)
	return resp.Parse(t2)
}

func innerSendContext(c *pemapi.Context) error {
	return innerSendMessage(c.Request, c.Response)
}
