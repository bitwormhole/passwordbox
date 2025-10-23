package weixin4pwbox
import (
    p3f194726f "github.com/bitwormhole/passwordbox/src/weixin/webapp"
     "github.com/starter-go/application"
)

// type p3f194726f.Demo1 in package:github.com/bitwormhole/passwordbox/src/weixin/webapp
//
// id:com-3f194726f1d7b174-webapp-Demo1
// class:
// alias:
// scope:singleton
//
type p3f194726f1_webapp_Demo1 struct {
}

func (inst* p3f194726f1_webapp_Demo1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3f194726f1d7b174-webapp-Demo1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3f194726f1_webapp_Demo1) new() any {
    return &p3f194726f.Demo1{}
}

func (inst* p3f194726f1_webapp_Demo1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3f194726f.Demo1)
	nop(ie, com)

	


    return nil
}


