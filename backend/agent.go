package backend

import (
	"github.com/bitwormhole/passwordbox/core/boot"
	"github.com/bitwormhole/passwordbox/core/components"
)

type agent struct {
	// ready bool

	backend *Backend

	context components.Context
}

var theBackendAgent agent

func (inst *agent) isReady() bool {
	return (inst.backend != nil) && (inst.context != nil)
}

func (inst *agent) innerTryLoad() *agent {

	if inst.isReady() {
		return inst
	}

	a2, err := inst.innerDoLoad()
	if err != nil {
		panic(err)
	}

	return a2
}

func (inst *agent) innerDoLoad() (*agent, error) {

	cfg1 := new(config)
	cfg2 := new(boot.Configuration)

	cfg1.configure(cfg2)
	cfg2.RegisterComponentFn(inst.registerComBackend)

	loader := new(boot.ContextLoader)
	ctx, err := loader.Load(cfg2)
	if err != nil {
		return nil, err
	}

	be := ctx.GetComponents().GetComponent("backend")
	inst.backend = be.(*Backend)
	inst.context = ctx
	return inst, nil
}

func (inst *agent) getBackend() *Backend {
	return inst.innerTryLoad().backend
}

func (inst *agent) getContext() components.Context {
	return inst.innerTryLoad().context
}

func (inst *agent) registerComBackend() *components.ComponentRegistration {

	r1 := new(components.ComponentRegistration)
	r1.ID = "backend"
	r1.OnNew = func() any {
		return new(Backend)
	}

	r1.OnLoad = func(l *components.Loading) {
		be := l.Component.(*Backend)
		be.foo = "bar"
	}

	return r1
}
