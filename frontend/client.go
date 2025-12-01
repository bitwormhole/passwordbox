package frontend

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/backend"
	"github.com/bitwormhole/passwordbox/core/pemapi"
)

type Client struct {
}

func (inst *Client) ExecuteMessage(request *pemapi.Request, response *pemapi.Response) error {

	if request == nil || response == nil {
		return fmt.Errorf("param (request|response) is nil")
	}

	codec := new(pemapi.CODEC)
	t1, err := codec.EncodeRequest(request)
	if err != nil {
		return err
	}

	t2 := inst.ExecuteText(t1)
	tmpResp, err := codec.DecodeResponse(t2)
	if err != nil {
		return err
	}

	*response = *tmpResp
	return nil
}

func (inst *Client) Execute(request *pemapi.Request, response *pemapi.Response) error {
	return inst.ExecuteMessage(request, response)
}

func (inst *Client) ExecuteContext(c *pemapi.Context) error {
	return inst.ExecuteMessage(c.Request, c.Response)
}

func (inst *Client) ExecuteString(request string) (response string) {
	api := backend.GetAPI()
	return api.ExecuteString(request)
}

func (inst *Client) ExecuteText(request pemapi.RequestText) (response pemapi.ResponseText) {
	str1 := request.String()
	str2 := inst.ExecuteString(str1)
	return pemapi.ResponseText(str2)
}

func NewClient() *Client {
	return new(Client)
}

func (inst *Client) _impl() pemapi.API {
	return inst
}
