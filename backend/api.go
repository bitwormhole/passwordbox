package backend

import (
	"net/http"

	"github.com/bitwormhole/passwordbox/core/pemapi"
)

func ExecuteString(request string) (response string) {
	t1 := pemapi.RequestText(request)
	t2 := innerHandleText(t1)
	return t2.String()
}

func innerHandleText(request pemapi.RequestText) (response pemapi.ResponseText) {

	m1 := new(pemapi.Request)
	m2 := new(pemapi.Response)

	err := m1.Parse(request)
	if err == nil {
		err = innerHandleMessage(m1, m2)
	}

	if err != nil {
		m2.Status.Code = http.StatusInternalServerError
		m2.SetHeader("error", err.Error())
	}

	return m2.Text()
}

func innerHandleMessage(request *pemapi.Request, response *pemapi.Response) error {
	ctx := new(pemapi.Context)
	ctx.Request = request
	ctx.Response = response
	return innerHandleContext(ctx)
}

func innerHandleContext(c *pemapi.Context) error {
	rt, err := GetDefaultRouter()
	if err != nil {
		return err
	}
	return rt.Handle(c)
}

func GetAPI() *pemapi.API {
	api := new(pemapi.API)
	api.StringHandler = ExecuteString
	api.MessageHandler = innerHandleMessage
	api.TextHandler = innerHandleText
	api.ContextHandler = innerHandleContext
	return api
}
