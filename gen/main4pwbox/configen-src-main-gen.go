package main4pwbox
import (
    p5ce388db2 "github.com/bitwormhole/passwordbox/app/classes/passwords"
    pc22dbb856 "github.com/bitwormhole/passwordbox/app/components/ipasswords"
    p9f2da412c "github.com/bitwormhole/passwordbox/app/web/controllers"
    pd1a916a20 "github.com/starter-go/libgin"
     "github.com/starter-go/application"
)

// type pc22dbb856.PasswordServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordServiceImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-Service
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordServiceImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordServiceImpl) new() any {
    return &pc22dbb856.PasswordServiceImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordServiceImpl)
	nop(ie, com)

	


    return nil
}



// type p9f2da412c.ExampleController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_ExampleController struct {
}

func (inst* p9f2da412c5_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_ExampleController) new() any {
    return &p9f2da412c.ExampleController{}
}

func (inst* p9f2da412c5_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p9f2da412c.PasswordController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-PasswordController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_PasswordController struct {
}

func (inst* p9f2da412c5_controllers_PasswordController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-PasswordController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_PasswordController) new() any {
    return &p9f2da412c.PasswordController{}
}

func (inst* p9f2da412c5_controllers_PasswordController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.PasswordController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_PasswordController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9f2da412c5_controllers_PasswordController) getService(ie application.InjectionExt)p5ce388db2.Service{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-Service").(p5ce388db2.Service)
}


