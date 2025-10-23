package controllers

import (
	"strconv"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/bitwormhole/passwordbox/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type AuthController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
}

func (inst *AuthController) _impl() libgin.Controller {
	return inst
}

func (inst *AuthController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *AuthController) route(rp libgin.RouterProxy) error {

	rp = rp.For("auth")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)

	return nil
}

func (inst *AuthController) handleGetOne(gc *gin.Context) {

	req := new(myAuthRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *AuthController) handleGetList(gc *gin.Context) {

	req := new(myAuthRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *AuthController) handlePutItem(gc *gin.Context) {

	req := new(myAuthRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myAuthRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *AuthController

	id    dxo.AuthID
	body1 vo.Authx
	body2 vo.Authx
}

func (inst *myAuthRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.AuthID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myAuthRequest) doGetList() error {

	it := &dto.Auth{}

	inst.body2.Items = []*dto.Auth{it, it, it}
	return nil
}

func (inst *myAuthRequest) doGetOne() error {

	it := &dto.Auth{}

	inst.body2.Items = []*dto.Auth{it}
	return nil
}

func (inst *myAuthRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Auth{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Auth{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
