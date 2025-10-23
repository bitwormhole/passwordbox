package controllers

import (
	"strconv"

	"github.com/bitwormhole/passwordbox/app/classes/hubproviders"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/bitwormhole/passwordbox/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

type ProviderController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender  libgin.Responder     //starter:inject("#")
	Service hubproviders.Service //starter:inject("#")
}

func (inst *ProviderController) _impl() libgin.Controller {
	return inst
}

func (inst *ProviderController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ProviderController) route(rp libgin.RouterProxy) error {

	rp = rp.For("providers")

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

	rp.PUT(":id", inst.handlePutItem)

	return nil
}

func (inst *ProviderController) handleGetOne(gc *gin.Context) {

	req := new(myProviderRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *ProviderController) handleGetList(gc *gin.Context) {

	req := new(myProviderRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *ProviderController) handlePutItem(gc *gin.Context) {

	req := new(myProviderRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myProviderRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *ProviderController

	id    dxo.ProviderID
	body1 vo.Providers
	body2 vo.Providers
}

func (inst *myProviderRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.ProviderID(num)
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

func (inst *myProviderRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myProviderRequest) doGetList() error {

	ctx := inst.context
	ser := inst.controller.Service

	list, err := ser.ListAll(ctx)
	if err != nil {
		return err
	}

	inst.body2.Items = list
	return nil
}

func (inst *myProviderRequest) doGetOne() error {

	it := &dto.Provider{}

	inst.body2.Items = []*dto.Provider{it}
	return nil
}

func (inst *myProviderRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Provider{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Provider{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
