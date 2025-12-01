package boot

import (
	"fmt"
	"time"

	"github.com/bitwormhole/passwordbox/core/components"
	"github.com/bitwormhole/passwordbox/core/errors"
	"github.com/bitwormhole/passwordbox/core/loggers"
)

type ContextLoader struct {
	context components.Context
	started bool
	stopped bool
	err     error
}

func (inst *ContextLoader) Load(cfg *Configuration) (components.Context, error) {

	ttl := cfg.Timeout
	if ttl < 1 {
		ttl = 15 * time.Second
	}

	go inst.run(cfg)

	ctx, err := inst.waitForContext(ttl)
	if ctx != nil {
		loggers.LogI("load_context: ok")
	}

	return ctx, err
}

func (inst *ContextLoader) waitForContext(ttl time.Duration) (components.Context, error) {
	const interval = time.Millisecond * 200
	for ttl > 0 {
		ttl -= interval
		ctx := inst.context
		if ctx != nil {
			return ctx, nil
		}
		if inst.stopped {
			break
		}
		time.Sleep(interval)
	}

	err := inst.err
	if err != nil {
		return nil, err
	}

	if ttl > 0 {
		ctx := inst.context
		if ctx != nil {
			return ctx, nil
		} else {
			return nil, fmt.Errorf("ContextLoader.waitForContext() : context is nil")
		}
	}

	return nil, fmt.Errorf("ContextLoader.waitForContext() : timeout")
}

func (inst *ContextLoader) run(cfg *Configuration) {

	defer func() {
		p := recover()
		err := errors.PanicToError(p)
		if err != nil {
			errors.HandleError(err)
			inst.err = err
		}
		inst.stopped = true
	}()
	inst.started = true

	cfg.RegisterComponentFn(inst.registerSelf)

	err := Run(cfg)
	if err == nil {
		return
	}
	inst.err = err
	// panic(err)
}

func (inst *ContextLoader) registerSelf() *components.ComponentRegistration {

	reg := new(components.ComponentRegistration)
	reg.ID = "boot.ContextLoader"

	reg.OnNew = func() any {
		return inst
	}

	reg.OnLoad = func(l *components.Loading) {
		inst.context = l.Context
	}

	return reg
}
