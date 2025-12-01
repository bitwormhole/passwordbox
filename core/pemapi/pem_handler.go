package pemapi

type Handler interface {
	Handle(c *Context) error
}

type HandlerExt interface {
	Handler

	ListRegistrations() []*HandlerRegistration
}

type HandlerFunc func(c *Context) error

type HandlerRegistration struct {
	Name     string
	Method   string
	Path     string
	Priority int
	Enabled  bool

	Handler Handler
	Filter  Filter
}

type HandlerRegistry interface {
	Register(info *HandlerRegistration)

	RegisterHandler(h HandlerExt)

	ListAll() []*HandlerRegistration
}

////////////////////////////////////////////////////////////////////////////////

type HandlerHolder struct {
	Key     HandlerKey
	Handler Handler
}

////////////////////////////////////////////////////////////////////////////////

type HandlerAdapter struct {
	Fn HandlerFunc
}

func (inst *HandlerAdapter) Handle(c *Context) error {
	return inst.Fn(c)
}

func (inst *HandlerAdapter) _impl() Handler {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func NewHandlerRegistry() HandlerRegistry {
	return new(innerHandlerRegistry)
}

////////////////////////////////////////////////////////////////////////////////

type innerHandlerRegistry struct {
	all []*HandlerRegistration
}

// ListAll implements HandlerRegistry.
func (i *innerHandlerRegistry) ListAll() []*HandlerRegistration {
	return i.all
}

func (i *innerHandlerRegistry) RegisterHandler(h HandlerExt) {
	if h == nil {
		return
	}
	list := h.ListRegistrations()
	i.all = append(i.all, list...)
}

func (i *innerHandlerRegistry) Register(info *HandlerRegistration) {
	if info == nil {
		return
	}
	i.all = append(i.all, info)
}

////////////////////////////////////////////////////////////////////////////////
