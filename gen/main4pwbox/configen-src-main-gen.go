package main4pwbox
import (
    pbdb8d38a2 "github.com/bitwormhole/passwordbox/app/classes/hubproviders"
    p5ce388db2 "github.com/bitwormhole/passwordbox/app/classes/passwords"
    p7c93e9365 "github.com/bitwormhole/passwordbox/app/classes/users"
    p916a30283 "github.com/bitwormhole/passwordbox/app/components/idb"
    p510a2f38b "github.com/bitwormhole/passwordbox/app/components/ihubproviders"
    pc22dbb856 "github.com/bitwormhole/passwordbox/app/components/ipasswords"
    pd01b6bc76 "github.com/bitwormhole/passwordbox/app/components/iusers"
    p2b6038bff "github.com/bitwormhole/passwordbox/app/data/database"
    p9f2da412c "github.com/bitwormhole/passwordbox/app/web/controllers"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p9621e8b71 "github.com/starter-go/security/random"
     "github.com/starter-go/application"
)

// type p916a30283.MyDataGroup in package:github.com/bitwormhole/passwordbox/app/components/idb
//
// id:com-916a30283a651046-idb-MyDataGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:
// scope:singleton
//
type p916a30283a_idb_MyDataGroup struct {
}

func (inst* p916a30283a_idb_MyDataGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-916a30283a651046-idb-MyDataGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p916a30283a_idb_MyDataGroup) new() any {
    return &p916a30283.MyDataGroup{}
}

func (inst* p916a30283a_idb_MyDataGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p916a30283.MyDataGroup)
	nop(ie, com)

	
    com.Alias = inst.getAlias(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.URI = inst.getURI(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*p916a30283a_idb_MyDataGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passwordbox.alias}")
}


func (inst*p916a30283a_idb_MyDataGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passwordbox.table-name-prefix}")
}


func (inst*p916a30283a_idb_MyDataGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passwordbox.datasource}")
}


func (inst*p916a30283a_idb_MyDataGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.passwordbox.uri}")
}


func (inst*p916a30283a_idb_MyDataGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.passwordbox.enabled}")
}



// type p916a30283.MyDatabaseAgent in package:github.com/bitwormhole/passwordbox/app/components/idb
//
// id:com-916a30283a651046-idb-MyDatabaseAgent
// class:
// alias:alias-2b6038bffa71faffd95c7cfc51d90a77-Agent
// scope:singleton
//
type p916a30283a_idb_MyDatabaseAgent struct {
}

func (inst* p916a30283a_idb_MyDatabaseAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-916a30283a651046-idb-MyDatabaseAgent"
	r.Classes = ""
	r.Aliases = "alias-2b6038bffa71faffd95c7cfc51d90a77-Agent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p916a30283a_idb_MyDatabaseAgent) new() any {
    return &p916a30283.MyDatabaseAgent{}
}

func (inst* p916a30283a_idb_MyDatabaseAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p916a30283.MyDatabaseAgent)
	nop(ie, com)

	
    com.DSM = inst.getDSM(ie)


    return nil
}


func (inst*p916a30283a_idb_MyDatabaseAgent) getDSM(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p510a2f38b.ProviderServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/ihubproviders
//
// id:com-510a2f38b08130d8-ihubproviders-ProviderServiceImpl
// class:
// alias:alias-bdb8d38a2d99a3c1502a66baf26b9cd6-Service
// scope:singleton
//
type p510a2f38b0_ihubproviders_ProviderServiceImpl struct {
}

func (inst* p510a2f38b0_ihubproviders_ProviderServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-510a2f38b08130d8-ihubproviders-ProviderServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-bdb8d38a2d99a3c1502a66baf26b9cd6-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p510a2f38b0_ihubproviders_ProviderServiceImpl) new() any {
    return &p510a2f38b.ProviderServiceImpl{}
}

func (inst* p510a2f38b0_ihubproviders_ProviderServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p510a2f38b.ProviderServiceImpl)
	nop(ie, com)

	


    return nil
}



// type pc22dbb856.PasswordBlockDaoImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordBlockDaoImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordBlockDaoImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordBlockDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockDaoImpl) new() any {
    return &pc22dbb856.PasswordBlockDaoImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordBlockDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordBlockDaoImpl) getAgent(ie application.InjectionExt)p2b6038bff.Agent{
    return ie.GetComponent("#alias-2b6038bffa71faffd95c7cfc51d90a77-Agent").(p2b6038bff.Agent)
}


func (inst*pc22dbb856b_ipasswords_PasswordBlockDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pc22dbb856.PasswordBlockServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordBlockServiceImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-BlockService
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordBlockServiceImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordBlockServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-BlockService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockServiceImpl) new() any {
    return &pc22dbb856.PasswordBlockServiceImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordBlockServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordBlockServiceImpl)
	nop(ie, com)

	
    com.DaoChains = inst.getDaoChains(ie)
    com.DaoBlocks = inst.getDaoBlocks(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordBlockServiceImpl) getDaoChains(ie application.InjectionExt)p5ce388db2.ChainDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO").(p5ce388db2.ChainDAO)
}


func (inst*pc22dbb856b_ipasswords_PasswordBlockServiceImpl) getDaoBlocks(ie application.InjectionExt)p5ce388db2.BlockDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO").(p5ce388db2.BlockDAO)
}



// type pc22dbb856.PasswordChainDaoImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordChainDaoImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordChainDaoImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordChainDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordChainDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordChainDaoImpl) new() any {
    return &pc22dbb856.PasswordChainDaoImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordChainDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordChainDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDService = inst.getUUIDService(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordChainDaoImpl) getAgent(ie application.InjectionExt)p2b6038bff.Agent{
    return ie.GetComponent("#alias-2b6038bffa71faffd95c7cfc51d90a77-Agent").(p2b6038bff.Agent)
}


func (inst*pc22dbb856b_ipasswords_PasswordChainDaoImpl) getUUIDService(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pc22dbb856.PasswordChainServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordChainServiceImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-ChainService
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordChainServiceImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordChainServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordChainServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-ChainService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordChainServiceImpl) new() any {
    return &pc22dbb856.PasswordChainServiceImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordChainServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordChainServiceImpl)
	nop(ie, com)

	
    com.DaoChains = inst.getDaoChains(ie)
    com.DaoBlocks = inst.getDaoBlocks(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordChainServiceImpl) getDaoChains(ie application.InjectionExt)p5ce388db2.ChainDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO").(p5ce388db2.ChainDAO)
}


func (inst*pc22dbb856b_ipasswords_PasswordChainServiceImpl) getDaoBlocks(ie application.InjectionExt)p5ce388db2.BlockDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO").(p5ce388db2.BlockDAO)
}



// type pc22dbb856.PasswordFastServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/ipasswords
//
// id:com-c22dbb856b626c34-ipasswords-PasswordFastServiceImpl
// class:
// alias:alias-5ce388db2e113c065a05f219a63b6c9e-FastService
// scope:singleton
//
type pc22dbb856b_ipasswords_PasswordFastServiceImpl struct {
}

func (inst* pc22dbb856b_ipasswords_PasswordFastServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c22dbb856b626c34-ipasswords-PasswordFastServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-5ce388db2e113c065a05f219a63b6c9e-FastService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc22dbb856b_ipasswords_PasswordFastServiceImpl) new() any {
    return &pc22dbb856.PasswordFastServiceImpl{}
}

func (inst* pc22dbb856b_ipasswords_PasswordFastServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc22dbb856.PasswordFastServiceImpl)
	nop(ie, com)

	
    com.DaoRefs = inst.getDaoRefs(ie)
    com.DaoBlocks = inst.getDaoBlocks(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordFastServiceImpl) getDaoRefs(ie application.InjectionExt)p5ce388db2.ChainDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO").(p5ce388db2.ChainDAO)
}


func (inst*pc22dbb856b_ipasswords_PasswordFastServiceImpl) getDaoBlocks(ie application.InjectionExt)p5ce388db2.BlockDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO").(p5ce388db2.BlockDAO)
}



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

	
    com.DaoChains = inst.getDaoChains(ie)
    com.DaoBlocks = inst.getDaoBlocks(ie)


    return nil
}


func (inst*pc22dbb856b_ipasswords_PasswordServiceImpl) getDaoChains(ie application.InjectionExt)p5ce388db2.ChainDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-ChainDAO").(p5ce388db2.ChainDAO)
}


func (inst*pc22dbb856b_ipasswords_PasswordServiceImpl) getDaoBlocks(ie application.InjectionExt)p5ce388db2.BlockDAO{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-BlockDAO").(p5ce388db2.BlockDAO)
}



// type pd01b6bc76.UserDaoImpl in package:github.com/bitwormhole/passwordbox/app/components/iusers
//
// id:com-d01b6bc76b7f6b88-iusers-UserDaoImpl
// class:
// alias:alias-7c93e93654a56c769229cefbf1acf168-DAO
// scope:singleton
//
type pd01b6bc76b_iusers_UserDaoImpl struct {
}

func (inst* pd01b6bc76b_iusers_UserDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d01b6bc76b7f6b88-iusers-UserDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7c93e93654a56c769229cefbf1acf168-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd01b6bc76b_iusers_UserDaoImpl) new() any {
    return &pd01b6bc76.UserDaoImpl{}
}

func (inst* pd01b6bc76b_iusers_UserDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd01b6bc76.UserDaoImpl)
	nop(ie, com)

	


    return nil
}



// type pd01b6bc76.UserServiceImpl in package:github.com/bitwormhole/passwordbox/app/components/iusers
//
// id:com-d01b6bc76b7f6b88-iusers-UserServiceImpl
// class:
// alias:alias-7c93e93654a56c769229cefbf1acf168-Service
// scope:singleton
//
type pd01b6bc76b_iusers_UserServiceImpl struct {
}

func (inst* pd01b6bc76b_iusers_UserServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d01b6bc76b7f6b88-iusers-UserServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7c93e93654a56c769229cefbf1acf168-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd01b6bc76b_iusers_UserServiceImpl) new() any {
    return &pd01b6bc76.UserServiceImpl{}
}

func (inst* pd01b6bc76b_iusers_UserServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd01b6bc76.UserServiceImpl)
	nop(ie, com)

	


    return nil
}



// type p9f2da412c.AuthController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-AuthController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_AuthController struct {
}

func (inst* p9f2da412c5_controllers_AuthController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-AuthController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_AuthController) new() any {
    return &p9f2da412c.AuthController{}
}

func (inst* p9f2da412c5_controllers_AuthController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.AuthController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_AuthController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
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



// type p9f2da412c.ProviderController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-ProviderController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_ProviderController struct {
}

func (inst* p9f2da412c5_controllers_ProviderController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-ProviderController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_ProviderController) new() any {
    return &p9f2da412c.ProviderController{}
}

func (inst* p9f2da412c5_controllers_ProviderController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.ProviderController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_ProviderController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9f2da412c5_controllers_ProviderController) getService(ie application.InjectionExt)pbdb8d38a2.Service{
    return ie.GetComponent("#alias-bdb8d38a2d99a3c1502a66baf26b9cd6-Service").(pbdb8d38a2.Service)
}



// type p9f2da412c.KeyPairController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-KeyPairController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_KeyPairController struct {
}

func (inst* p9f2da412c5_controllers_KeyPairController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-KeyPairController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_KeyPairController) new() any {
    return &p9f2da412c.KeyPairController{}
}

func (inst* p9f2da412c5_controllers_KeyPairController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.KeyPairController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_KeyPairController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
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
    com.FastService = inst.getFastService(ie)
    com.Chains = inst.getChains(ie)
    com.Blocks = inst.getBlocks(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_PasswordController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9f2da412c5_controllers_PasswordController) getService(ie application.InjectionExt)p5ce388db2.Service{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-Service").(p5ce388db2.Service)
}


func (inst*p9f2da412c5_controllers_PasswordController) getFastService(ie application.InjectionExt)p5ce388db2.FastService{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-FastService").(p5ce388db2.FastService)
}


func (inst*p9f2da412c5_controllers_PasswordController) getChains(ie application.InjectionExt)p5ce388db2.ChainService{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-ChainService").(p5ce388db2.ChainService)
}


func (inst*p9f2da412c5_controllers_PasswordController) getBlocks(ie application.InjectionExt)p5ce388db2.BlockService{
    return ie.GetComponent("#alias-5ce388db2e113c065a05f219a63b6c9e-BlockService").(p5ce388db2.BlockService)
}



// type p9f2da412c.UserController in package:github.com/bitwormhole/passwordbox/app/web/controllers
//
// id:com-9f2da412c56a3a0d-controllers-UserController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9f2da412c5_controllers_UserController struct {
}

func (inst* p9f2da412c5_controllers_UserController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9f2da412c56a3a0d-controllers-UserController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9f2da412c5_controllers_UserController) new() any {
    return &p9f2da412c.UserController{}
}

func (inst* p9f2da412c5_controllers_UserController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9f2da412c.UserController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p9f2da412c5_controllers_UserController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9f2da412c5_controllers_UserController) getService(ie application.InjectionExt)p7c93e9365.Service{
    return ie.GetComponent("#alias-7c93e93654a56c769229cefbf1acf168-Service").(p7c93e9365.Service)
}


