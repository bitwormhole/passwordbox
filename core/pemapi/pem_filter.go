package pemapi

type Filter interface {
	Pass(c *Context, next FilterChain) error
}

type FilterChain interface {
	Pass(c *Context) error
}

type FilterFunc func(c *Context, next FilterChain) error

////////////////////////////////////////////////////////////////////////////////

type FilterChainBuilder struct {
	items []*HandlerRegistration
}

func (inst *FilterChainBuilder) Add(item ...*HandlerRegistration) {
	if item == nil {
		return
	}
	inst.items = append(inst.items, item...)
}

func (inst *FilterChainBuilder) Build() FilterChain {

	items := inst.items
	end := new(innerFilterChainEnding)
	ptr := end._impl()

	for _, item := range items {
		node := new(innerFilterChainNode)
		node.filter = item.Filter
		node.next = ptr
		ptr = node
	}

	return ptr
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterChainEnding struct {
}

// Pass implements FilterChain.
func (inst *innerFilterChainEnding) Pass(c *Context) error {
	return nil // nop
}

func (inst *innerFilterChainEnding) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterChainNode struct {
	filter Filter
	next   FilterChain
}

func (inst *innerFilterChainNode) Pass(c *Context) error {
	n := inst.next
	f := inst.filter
	return f.Pass(c, n)
}

func (inst *innerFilterChainNode) _impl() FilterChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type FilterAdapter struct {
	Fn FilterFunc
}

func (inst *FilterAdapter) Pass(c *Context, next FilterChain) error {
	return inst.Fn(c, next)
}

func (inst *FilterAdapter) _impl() Filter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
