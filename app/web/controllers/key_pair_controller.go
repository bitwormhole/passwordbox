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

type KeyPairController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")
}

func (inst *KeyPairController) _impl() libgin.Controller {
	return inst
}

func (inst *KeyPairController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *KeyPairController) route(rp libgin.RouterProxy) error {

	rp = rp.For("key-pairs")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)

	return nil
}

func (inst *KeyPairController) handleGetOne(gc *gin.Context) {

	req := new(myKeyPairRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *KeyPairController) handleGetList(gc *gin.Context) {

	req := new(myKeyPairRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *KeyPairController) handlePutItem(gc *gin.Context) {

	req := new(myKeyPairRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myKeyPairRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *KeyPairController

	id    dxo.KeyPairID
	body1 vo.KeyPairs
	body2 vo.KeyPairs
}

func (inst *myKeyPairRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.KeyPairID(num)
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

func (inst *myKeyPairRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myKeyPairRequest) doGetList() error {

	it := &dto.KeyPair{}

	inst.body2.Items = []*dto.KeyPair{it, it, it}
	return nil
}

func (inst *myKeyPairRequest) doGetOne() error {

	it := &dto.KeyPair{}

	inst.body2.Items = []*dto.KeyPair{it}
	return nil
}

func (inst *myKeyPairRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.KeyPair{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.KeyPair{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
