package main

import (
	"fmt"
	"net/http"

	"github.com/bitwormhole/passwordbox/core/pemapi"
	"github.com/bitwormhole/passwordbox/frontend"
)

func main() {

	fmt.Println("hello, pbox main")

	api := frontend.GetAPI()
	ctx := new(pemapi.Context)

	ctx.Request = new(pemapi.Request)
	ctx.Response = new(pemapi.Response)

	ctx.Request.Method = http.MethodGet
	ctx.Request.SetLocation("pemapi://user@host/api/mock/demo1")

	err := api.ExecuteContext(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
