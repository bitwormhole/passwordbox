package boot

import (
	"time"

	"github.com/bitwormhole/passwordbox/core/components"
)

type Configuration struct {
	Arguments   []string
	Components  components.Registry
	Environment map[string]string
	Properties  map[string]string
	Timeout     time.Duration
}

func (inst *Configuration) innerInit() *Configuration {

	if inst.Arguments == nil {
		inst.Arguments = make([]string, 0)
	}

	if inst.Components == nil {
		inst.Components = components.NewRegistry()
	}

	if inst.Environment == nil {
		inst.Environment = make(map[string]string)
	}

	if inst.Properties == nil {
		inst.Properties = make(map[string]string)
	}

	return inst
}

func (inst *Configuration) SetProperty(name, value string) {
	in := inst.innerInit()
	table := in.Properties
	table[name] = value
}

func (inst *Configuration) SetEnv(name, value string) {
	in := inst.innerInit()
	table := in.Environment
	table[name] = value
}

func (inst *Configuration) RegisterComponent(cr *components.ComponentRegistration) {
	in := inst.innerInit()
	reg := in.Components
	reg.Register(cr)
}

func (inst *Configuration) RegisterComponentFn(fn components.ComponentRegistryFunc) {
	in := inst.innerInit()
	reg := in.Components
	reg.RegisterFn(fn)
}
