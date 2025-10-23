package controllers

import (
	"strconv"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/bitwormhole/passwordbox/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type PasswordController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

	Service passwords.Service //starter:inject("#")
}

func (inst *PasswordController) _impl() libgin.Controller {
	return inst
}

func (inst *PasswordController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *PasswordController) route(rp libgin.RouterProxy) error {

	rp = rp.For("passwords")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutOne)
	rp.DELETE(":id", inst.handleDeleteOne)

	rp.POST("", inst.handlePostOne)
	rp.POST("do/init-new-password", inst.handlePostExample)
	rp.POST(":id/create-new-revision", inst.handlePostExample)
	rp.POST(":id/apply", inst.handlePostExample)

	return nil
}

func (inst *PasswordController) handleGetOne(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = false

	req.execute(req.doFindItem)
}

func (inst *PasswordController) handleGetList(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doQueryItems)
}

func (inst *PasswordController) handlePostOne(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doInsertItem)
}

func (inst *PasswordController) handlePutOne(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doUpdateItem)
}

func (inst *PasswordController) handleDeleteOne(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = false

	req.execute(req.doRemoveItem)
}

func (inst *PasswordController) handlePostExample(gc *gin.Context) {

	req := new(myPasswordRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = true

	req.execute(req.doExample)
}

////////////////////////////////////////////////////////////////////////////////

type myPasswordRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *PasswordController

	id    dxo.PasswordID
	body1 vo.Passwords
	body2 vo.Passwords
}

func (inst *myPasswordRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.PasswordID(num)
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

func (inst *myPasswordRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myPasswordRequest) doQueryItems() error {

	ctx := inst.context
	ser := inst.controller.Service
	q := new(passwords.Query)

	// todo: fill query

	list, err := ser.Query(ctx, q)
	if err != nil {
		return err
	}

	inst.body2.Items = list
	return nil
}

func (inst *myPasswordRequest) doFindItem() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.id

	item, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Password{item}
	return nil
}

func (inst *myPasswordRequest) doInsertItem() error {

	ctx := inst.context
	ser := inst.controller.Service
	it1 := inst.body1.Items[0]

	it2, err := ser.Insert(ctx, it1)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Password{it2}
	return nil
}

func (inst *myPasswordRequest) doUpdateItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Password{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Password{it1, it2}
	return nil
}

func (inst *myPasswordRequest) doRemoveItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Password{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Password{it1, it2}
	return nil
}

func (inst *myPasswordRequest) doExample() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Password{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Password{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
