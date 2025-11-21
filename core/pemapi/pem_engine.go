package pemapi

import (
	"fmt"
	"strings"
	"sync"
)

////////////////////////////////////////////////////////////////////////////////

type Router struct {
	config   *RouterConfig
	handlers map[HandlerKey]*HandlerRegistration
	mutex    sync.Mutex
}

// ListRegistrations implements Handler.
func (inst *Router) ListRegistrations() []*HandlerRegistration {
	return nil
}

// Handle implements Handler.
func (inst *Router) Handle(c *Context) error {

	hkb := new(HandlerKeyBuilder)
	hkb.Method = c.Request.Method
	hkb.Path = c.Request.Path

	next, err := inst.findNextHandler(hkb.Key())
	if err != nil {
		return err
	}
	return next(c)
}

func (inst *Router) getHandler() Handler {
	return inst
}

func (inst *Router) findNextHandler(key HandlerKey) (HandlerFunc, error) {

	mtx := &inst.mutex
	mtx.Lock()
	defer mtx.Unlock()

	all := inst.handlers
	if all != nil {
		next := all[key]
		if next != nil {
			return next.Handler, nil
		}
	}

	return nil, fmt.Errorf("pemapi.Router: no handler for api [%s]", key)
}

////////////////////////////////////////////////////////////////////////////////

type HandlerKey string

type HandlerKeyBuilder struct {
	Method string
	Path   string
}

func (inst *HandlerKeyBuilder) Key() HandlerKey {
	s1 := strings.ToUpper(inst.Method)
	s2 := strings.TrimSpace(inst.Path)
	str := s1 + ":" + s2
	return HandlerKey(str)
}

////////////////////////////////////////////////////////////////////////////////

type RouterConfig struct {
	Handlers HandlerRegistry

	Loader RouterLoader
}

////////////////////////////////////////////////////////////////////////////////

type RouterLoader interface {
	Load(config *RouterConfig) (*Router, error)
}

func NewRouterLoader() RouterLoader {
	loader := new(innerRouterLoader)
	return loader
}

////////////////////////////////////////////////////////////////////////////////

type innerRouterLoader struct{}

// Load implements RouterLoader.
func (inst *innerRouterLoader) Load(config *RouterConfig) (*Router, error) {

	rt := new(Router)
	rt.handlers = make(map[HandlerKey]*HandlerRegistration)
	rt.config = config

	src := config.Handlers.ListAll()
	hkb := new(HandlerKeyBuilder)

	for _, item := range src {
		if !inst.innerIsHandlerReady(item) {
			continue
		}
		hkb.Path = item.Path
		hkb.Method = item.Method
		key := hkb.Key()
		rt.handlers[key] = item
	}

	return rt, nil
}

func (i *innerRouterLoader) innerIsHandlerReady(item *HandlerRegistration) bool {

	if item == nil {
		return false
	}

	if item.Path == "" {
		return false
	}
	if item.Handler == nil {
		return false
	}
	if !item.Enabled {
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////
