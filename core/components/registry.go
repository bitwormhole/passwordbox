package components

////////////////////////////////////////////////////////////////////////////////

// 表示组件的注册信息
type ComponentRegistration struct {
	ID      ID
	Scope   Scope
	Classes []Class
	Aliases []Alias

	OnNew  OnNewFunc
	OnLoad OnLoadFunc
}

////////////////////////////////////////////////////////////////////////////////

type Registry struct {
	list []*ComponentRegistration
}

func (inst *Registry) Register(cr *ComponentRegistration) {
	inst.list = append(inst.list, cr)
}

func (inst *Registry) RegisterFn(fn OnRegisterFunc) {
	cr := fn()
	inst.Register(cr)
}

func (inst *Registry) ListAll() []*ComponentRegistration {
	return inst.list
}

////////////////////////////////////////////////////////////////////////////////
