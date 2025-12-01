package boot

import "github.com/bitwormhole/passwordbox/core/components"

func Run(cfg *Configuration) error {
	r := new(runner)
	err := r.init(cfg)
	if err != nil {
		return err
	}
	return r.run()
}

////////////////////////////////////////////////////////////////////////////////

type runner struct {
	config     *Configuration
	ctxBuilder *components.ContextBuilder
	context    components.Context
	comLoader  components.Loader
}

func (inst *runner) init(cfg *Configuration) error {

	inst.config = cfg

	return nil
}

func (inst *runner) run() error {

	steps := make([]func() error, 0)

	steps = append(steps, inst.runStep2makeContextBuilder)
	steps = append(steps, inst.runStep2makeContext)

	steps = append(steps, inst.runStep2makeSingletonComponents)
	steps = append(steps, inst.runStep2wire)
	steps = append(steps, inst.runStep2applyLifeCycle)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *runner) runStep2makeContextBuilder() error {
	inst.ctxBuilder = new(components.ContextBuilder)
	return nil
}

func (inst *runner) configureArguments(src *Configuration, dst *components.ContextBuilder) {
	dst.SetArguments(src.Arguments)
}

func (inst *runner) configureEnvironment(src *Configuration, dst *components.ContextBuilder) {
	all := src.Environment
	for name, value := range all {
		dst.SetEnv(name, value)
	}
}

func (inst *runner) configureProperties(src *Configuration, dst *components.ContextBuilder) {
	all := src.Properties
	for name, value := range all {
		dst.SetProperty(name, value)
	}
}

func (inst *runner) configureComponents(src *Configuration, dst *components.ContextBuilder) {
	all := src.Components.ListAll()
	for _, cr := range all {
		dst.RegisterComponent(cr)
	}
}

func (inst *runner) runStep2makeContext() error {

	cfg := inst.config
	cb := inst.ctxBuilder

	inst.configureArguments(cfg, cb)
	inst.configureEnvironment(cfg, cb)
	inst.configureProperties(cfg, cb)
	inst.configureComponents(cfg, cb)

	ctx, err := cb.Build()
	if err != nil {
		return err
	}
	inst.context = ctx
	inst.comLoader = ctx.GetComponents().NewLoader()
	return nil
}

func (inst *runner) runStep2makeSingletonComponents() error {

	src := inst.context.GetComponents().GetRegistry().ListAll()
	loader := inst.comLoader

	for _, r1 := range src {
		id := r1.ID
		loader.GetComponentByID(id)
	}

	return nil
}

func (inst *runner) runStep2wire() error {
	loader := inst.comLoader
	return loader.Load()
}

func (inst *runner) runStep2applyLifeCycle() error {
	return nil
}
