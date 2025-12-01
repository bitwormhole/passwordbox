package frontend

import (
	"net/http"
	"testing"

	"github.com/bitwormhole/passwordbox/core/pemapi"
)

func TestPEMAPI(t *testing.T) {

	api := GetAPI()
	ctx := new(pemapi.Context)
	req := new(pemapi.Request)
	resp := new(pemapi.Response)

	ctx.Request = req
	ctx.Response = resp

	req.Method = http.MethodGet
	req.SetLocation("pemapi://user@host/test/mock-1")
	req.SetHeader("x-foo", "a header named 'foo'")
	req.SetHeader("x-bar", "a header named 'bar'")
	req.Query = map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	req.Params = map[string]string{
		"x": "7",
		"y": "8",
		"z": "9",
	}

	content := "hello, PEM-API"
	req.Bytes = []byte(content)

	err := api.ExecuteContext(ctx)
	if err == nil {
		t.Log("OK")
	} else {
		t.Error(err)
	}

	t.Log("Done.")
}
