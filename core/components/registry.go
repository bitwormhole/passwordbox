package components

import "github.com/bitwormhole/passwordbox/core/loggers"

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

func (inst *ComponentRegistration) SetClasses(cl ...Class) *ComponentRegistration {
	inst.Classes = cl
	return inst
}

func (inst *ComponentRegistration) SetAliases(al ...Alias) *ComponentRegistration {
	inst.Aliases = al
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type Registry interface {
	Register(cr *ComponentRegistration)

	RegisterFn(fn ComponentRegistryFunc)

	ListAll() []*ComponentRegistration

	ListIDs() []ID

	GetByID(id ID) *ComponentRegistration

	// 去掉重复的项
	Dedup() Registry
}

////////////////////////////////////////////////////////////////////////////////

type innerRegistryCache struct {
	ids   []ID
	all   []*ComponentRegistration
	table map[ID]*ComponentRegistration
}

func (inst *innerRegistryCache) init() {
	inst.all = make([]*ComponentRegistration, 0)
	inst.ids = make([]ID, 0)
	inst.table = make(map[ID]*ComponentRegistration)
}

////////////////////////////////////////////////////////////////////////////////

type innerRegistry struct {
	list  []*ComponentRegistration
	cache *innerRegistryCache
}

// Dedup implements Registry.
func (inst *innerRegistry) Dedup() Registry {

	src := inst.list
	dst := new(innerRegistry)
	tmp := make(map[ID]*ComponentRegistration)

	for _, item := range src {
		id := item.ID
		older := tmp[id]
		if older != nil {
			loggers.LogW("dedup component, id:%s", id)
		}
		tmp[id] = item
	}

	for _, item := range tmp {
		dst.list = append(dst.list, item)
	}

	return dst
}

func (inst *innerRegistry) innerGetCache() *innerRegistryCache {

	src := inst.list
	dst := new(innerRegistryCache)
	dst.init()

	for _, item := range src {
		id := item.ID
		dst.table[id] = item
		dst.all = append(dst.all, item)
		dst.ids = append(dst.ids, id)
	}

	return dst
}

// GetByID implements Registry.
func (inst *innerRegistry) GetByID(id ID) *ComponentRegistration {
	cache := inst.innerGetCache()
	return cache.table[id]
}

// ListIDs implements Registry.
func (inst *innerRegistry) ListIDs() []ID {
	cache := inst.innerGetCache()
	return cache.ids
}

func (inst *innerRegistry) init() Registry {
	inst.list = make([]*ComponentRegistration, 0)
	return inst
}

func (inst *innerRegistry) Register(cr *ComponentRegistration) {
	inst.list = append(inst.list, cr)
	inst.cache = nil
}

func (inst *innerRegistry) RegisterFn(fn ComponentRegistryFunc) {
	cr := fn()
	inst.Register(cr)
	inst.cache = nil
}

func (inst *innerRegistry) ListAll() []*ComponentRegistration {
	cache := inst.innerGetCache()
	return cache.all
}

////////////////////////////////////////////////////////////////////////////////

func NewRegistry() Registry {
	return new(innerRegistry)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
