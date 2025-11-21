package pemapi

type Handler interface {
	Handle(c *Context) error

	ListRegistrations() []*HandlerRegistration
}

type HandlerRegistration struct {
	Name     string
	Method   string
	Path     string
	Handler  HandlerFunc
	Priority int
	Enabled  bool
}

type HandlerRegistry interface {
	Register(info *HandlerRegistration)

	RegisterHandler(h Handler)

	ListAll() []*HandlerRegistration
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

func (i *innerHandlerRegistry) RegisterHandler(h Handler) {
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
