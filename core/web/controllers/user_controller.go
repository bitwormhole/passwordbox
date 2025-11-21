package controllers

import (
	"strconv"

	"github.com/bitwormhole/passwordbox/app/classes/users"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/bitwormhole/passwordbox/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type UserController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder //starter:inject("#")
	Service users.Service    //starter:inject("#")

}

func (inst *UserController) _impl() libgin.Controller {
	return inst
}

func (inst *UserController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *UserController) route(rp libgin.RouterProxy) error {

	rp = rp.For("users")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)

	return nil
}

func (inst *UserController) handleGetOne(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *UserController) handleGetList(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *UserController) handlePutItem(gc *gin.Context) {

	req := new(myUserRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myUserRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *UserController

	id    dxo.UserID
	body1 vo.Users
	body2 vo.Users
}

func (inst *myUserRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.UserID(num)
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

func (inst *myUserRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myUserRequest) doGetList() error {

	it := &dto.User{}

	inst.body2.Items = []*dto.User{it, it, it}
	return nil
}

func (inst *myUserRequest) doGetOne() error {

	it := &dto.User{}

	inst.body2.Items = []*dto.User{it}
	return nil
}

func (inst *myUserRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.User{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.User{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
