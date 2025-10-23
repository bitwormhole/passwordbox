package hub4pwbox
import (
    pa9a43e1d7 "github.com/bitwormhole/passwordbox/src/hub/webapp"
     "github.com/starter-go/application"
)

// type pa9a43e1d7.Demo1 in package:github.com/bitwormhole/passwordbox/src/hub/webapp
//
// id:com-a9a43e1d77a98afd-webapp-Demo1
// class:
// alias:
// scope:singleton
//
type pa9a43e1d77_webapp_Demo1 struct {
}

func (inst* pa9a43e1d77_webapp_Demo1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a9a43e1d77a98afd-webapp-Demo1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa9a43e1d77_webapp_Demo1) new() any {
    return &pa9a43e1d7.Demo1{}
}

func (inst* pa9a43e1d77_webapp_Demo1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa9a43e1d7.Demo1)
	nop(ie, com)

	


    return nil
}


