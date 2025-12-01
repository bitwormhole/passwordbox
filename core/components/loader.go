package components

import "fmt"

// Loader 表示组件的加载器
type Loader interface {
	GetComponentByID(id ID) any

	ListComponentsByClass(cls Class) []any

	ListAllComponents() []any

	ListAllComponentIDs() []ID

	Load() error
}

type Loading struct {
	Context   Context
	Loader    Loader
	Component any // the Component-instance
	ID        ID
}

// 组件的工厂函数
type OnNewFunc func() any

// 组件的加载函数
type OnLoadFunc func(l *Loading)

// 组件的注册函数
type ComponentRegistryFunc func() *ComponentRegistration

////////////////////////////////////////////////////////////////////////////////

type innerComponentLoader struct {
	context *innerComponentContext
	cache   *innerComponentCache
}

func (inst *innerComponentLoader) init(ctx *innerComponentContext) {
	inst.context = ctx
	inst.cache = new(innerComponentCache)
	inst.cache.init(ctx)
}

// ListAllComponentIDs implements Loader.
func (inst *innerComponentLoader) ListAllComponentIDs() []ID {
	src := inst.context.registry.ListAll()
	dst := make([]ID, 0)
	for _, reg := range src {
		dst = append(dst, reg.ID)
	}
	return dst
}

// GetComponentByID implements Loader.
func (inst *innerComponentLoader) GetComponentByID(id ID) any {
	h, err := inst.cache.fetchByID(id)
	if err != nil {
		panic(err)
	}
	return h.getInstance()
}

// ListAllComponents implements Loader.
func (inst *innerComponentLoader) ListAllComponents() []any {
	ids := inst.ListAllComponentIDs()
	dst := make([]any, 0)
	for _, id := range ids {
		com := inst.GetComponentByID(id)
		dst = append(dst, com)
	}
	return dst
}

// ListComponentsByClass implements Loader.
func (inst *innerComponentLoader) ListComponentsByClass(cls Class) []any {
	dst := make([]any, 0)
	src := inst.context.indexer.selectList(cls.Selector())
	for _, item := range src {
		com := item.getInstance()
		dst = append(dst, com)
	}
	return dst
}

func (inst *innerComponentLoader) Load() error {
	for ttl := 30; ttl > 0; ttl-- {
		count, err := inst.innerTryLoadOnce()
		if err != nil {
			return err
		}
		if count == 0 {
			return nil
		}
	}
	return fmt.Errorf("innerComponentLoader.Load() : timeout")
}

func (inst *innerComponentLoader) innerTryLoadOnce() (count int, err error) {

	count = 0
	table := inst.cache.items
	ids := make([]ID, 0)

	for _, holder := range table {
		ids = append(ids, holder.id)
	}

	for _, id := range ids {
		hdr2 := table[id]
		if hdr2.loaded {
			continue
		}
		err2 := inst.innerTryLoadItem(hdr2)
		if err2 != nil {
			return 0, err2
		}
		count++
	}

	return
}

func (inst *innerComponentLoader) innerTryLoadItem(h *innerComponentHolder) error {

	info := h.info
	fn := info.OnLoad
	com := h.getInstance()
	loading := new(Loading)

	loading.Component = com
	loading.Context = inst.context
	loading.Loader = inst
	loading.ID = h.id
	h.loader = inst

	fn(loading)

	h.loaded = true
	return nil
}

func (inst *innerComponentLoader) toLoader() Loader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
