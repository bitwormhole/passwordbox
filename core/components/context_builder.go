package components

import "github.com/bitwormhole/passwordbox/core/data/properties"

////////////////////////////////////////////////////////////////////////////////

type innerContextBuilderData struct {
	registry    Registry
	args        []string
	attributes  properties.Table
	properties  properties.Table
	environment properties.Table
	factory     ContextFactory
}

////////////////////////////////////////////////////////////////////////////////

type ContextBuilder struct {
	data *innerContextBuilderData
}

func (inst *ContextBuilder) innerGetData() *innerContextBuilderData {

	data := inst.data
	if data != nil {
		return data
	}

	data = new(innerContextBuilderData)

	data.factory = new(innerContextFactory)
	data.properties = properties.NewTable()
	data.registry = new(innerRegistry)
	data.attributes = properties.NewTable()
	data.environment = properties.NewTable()
	data.args = []string{}

	inst.data = data
	return data
}

func (inst *ContextBuilder) Build() (Context, error) {
	data := inst.innerGetData()

	data.registry = data.registry.Dedup()
	factory := data.factory

	return factory.Create(inst)
}

func (inst *ContextBuilder) SetArguments(args []string) *ContextBuilder {

	if args == nil {
		return inst
	}

	data := inst.innerGetData()
	data.args = args
	return inst
}

func (inst *ContextBuilder) SetAttribute(name string, value string) *ContextBuilder {
	data := inst.innerGetData()
	data.attributes.Put(name, value)
	return inst
}

func (inst *ContextBuilder) SetEnv(name string, value string) *ContextBuilder {
	data := inst.innerGetData()
	data.environment.Put(name, value)
	return inst
}

func (inst *ContextBuilder) SetProperty(name string, value string) *ContextBuilder {
	data := inst.innerGetData()
	data.properties.Put(name, value)
	return inst
}

func (inst *ContextBuilder) RegisterComponent(com *ComponentRegistration) *ContextBuilder {
	data := inst.innerGetData()
	data.registry.Register(com)
	return inst
}

func (inst *ContextBuilder) RegisterComponentFn(fn ComponentRegistryFunc) *ContextBuilder {
	data := inst.innerGetData()
	data.registry.RegisterFn(fn)
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
