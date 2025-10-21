package test4pwbox
import (
    pbdc24cec7 "github.com/bitwormhole/passwordbox/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type pbdc24cec7.ExampleCase in package:github.com/bitwormhole/passwordbox/src/test/golang/unittestcases
//
// id:com-bdc24cec72516d30-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type pbdc24cec72_unittestcases_ExampleCase struct {
}

func (inst* pbdc24cec72_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bdc24cec72516d30-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbdc24cec72_unittestcases_ExampleCase) new() any {
    return &pbdc24cec7.ExampleCase{}
}

func (inst* pbdc24cec72_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbdc24cec7.ExampleCase)
	nop(ie, com)

	


    return nil
}


