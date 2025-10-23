package main4pwbox

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p510a2f38b0_ihubproviders_ProviderServiceImpl{})
    inst.register(&p916a30283a_idb_MyDataGroup{})
    inst.register(&p916a30283a_idb_MyDatabaseAgent{})
    inst.register(&p9f2da412c5_controllers_AuthController{})
    inst.register(&p9f2da412c5_controllers_ExampleController{})
    inst.register(&p9f2da412c5_controllers_KeyPairController{})
    inst.register(&p9f2da412c5_controllers_PasswordController{})
    inst.register(&p9f2da412c5_controllers_ProviderController{})
    inst.register(&p9f2da412c5_controllers_UserController{})
    inst.register(&pc22dbb856b_ipasswords_PasswordDaoImpl{})
    inst.register(&pc22dbb856b_ipasswords_PasswordServiceImpl{})
    inst.register(&pd01b6bc76b_iusers_UserDaoImpl{})
    inst.register(&pd01b6bc76b_iusers_UserServiceImpl{})


    return nil
}
