package components

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/passwordbox/core/data/properties"
)

type ID string
type Class string
type Alias string
type Scope int
type Selector string
type Name string // name = [id|alias]

const (
	ScopeSingleton Scope = 1
	ScopePrototype Scope = 2
)

////////////////////////////////////////////////////////////////////////////////
// id

func (n ID) Selector() Selector {
	return Selector("#" + n)
}

////////////////////////////////////////////////////////////////////////////////
// class

func (n Class) Selector() Selector {
	return Selector("." + n)
}

////////////////////////////////////////////////////////////////////////////////
// alias

func (n Alias) Selector() Selector {
	return Selector("#" + n)
}

////////////////////////////////////////////////////////////////////////////////
// Selector

func (n Selector) String() string {
	return string(n)
}

////////////////////////////////////////////////////////////////////////////////

type ComponentManager interface {
	GetComponent(id ID) any

	TryGetComponent(id ID) (any, error)

	ListComponents(cl Class) []any

	ListIDs() []ID

	NewLoader() Loader

	GetRegistry() Registry
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentManager struct {
	context *innerComponentContext
}

// GetRegistry implements ComponentManager.
func (inst *innerComponentManager) GetRegistry() Registry {
	return inst.context.registry
}

// ListComponents implements ComponentManager.
func (inst *innerComponentManager) ListComponents(cl Class) []any {
	src := inst.context.indexer.selectList(cl.Selector())
	dst := make([]any, 0)
	for _, holder := range src {
		com := holder.getInstance()
		dst = append(dst, com)
	}
	return dst
}

// ListIDs implements ComponentManager.
func (inst *innerComponentManager) ListIDs() []ID {
	return inst.context.registry.ListIDs()
}

// NewLoader implements ComponentManager.
func (inst *innerComponentManager) NewLoader() Loader {

	loader := new(innerComponentLoader)
	loader.init(inst.context)

	loader.cache.parent = inst.context.cache

	return loader
}

func (inst *innerComponentManager) init(ctx *innerComponentContext) {
	inst.context = ctx
}

func (inst *innerComponentManager) GetComponent(id ID) any {
	com, err := inst.TryGetComponent(id)
	if err != nil {
		panic(err)
	}
	return com
}

// TryGetComponent implements ComponentManager.
func (inst *innerComponentManager) TryGetComponent(id ID) (any, error) {

	h, err := inst.context.cache.fetchByID(id)
	if err != nil {
		return nil, err
	}

	if h.loaded && h.instance != nil {
		return h.instance, nil
	}

	// do load
	loader := inst.NewLoader()
	com := loader.GetComponentByID(id)
	err = loader.Load()
	if err != nil {
		return nil, err
	}

	return com, nil
}

func (inst *innerComponentManager) facade() ComponentManager {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentCache struct {
	parent  *innerComponentCache
	context *innerComponentContext
	items   map[ID]*innerComponentHolder
}

func (inst *innerComponentCache) init(ctx *innerComponentContext) {
	inst.context = ctx
	inst.items = make(map[ID]*innerComponentHolder)
}

func (inst *innerComponentCache) fetchByID(id ID) (*innerComponentHolder, error) {

	older := inst.items[id]
	item := older
	parent := inst.parent

	if older != nil {
		return older, nil
	}

	if parent != nil {
		item2, err := parent.fetchByID(id)
		if err != nil {
			return nil, err
		}
		item = item2.tryClone()
	} else {
		item2, err := inst.loadNewItem(id)
		if err != nil {
			return nil, err
		}
		item = item2
	}

	inst.items[id] = item
	return item, nil
}

func (inst *innerComponentCache) loadNewItem(id ID) (*innerComponentHolder, error) {

	info := inst.context.registry.GetByID(id)
	if info == nil {
		return nil, fmt.Errorf("no component with id: [%s]", id)
	}

	holder := new(innerComponentHolder)
	holder.id = id
	holder.info = info
	holder.instance = nil
	holder.loaded = false
	holder.scope = info.Scope

	return holder, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentHolderChain struct {
	sel    Selector
	holder *innerComponentHolder
	next   *innerComponentHolderChain
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentHolder struct {
	id       ID
	scope    Scope
	loaded   bool
	loader   Loader
	context  *innerComponentContext
	info     *ComponentRegistration
	instance any
}

func (inst *innerComponentHolder) makeNewInstance() any {
	fn := inst.info.OnNew
	return fn()
}

func (inst *innerComponentHolder) getInstance() any {
	i := inst.instance
	if i == nil {
		i = inst.makeNewInstance()
		inst.instance = i
	}
	return i
}

func (inst *innerComponentHolder) innerDoRealClone() *innerComponentHolder {

	src := inst
	dst := new(innerComponentHolder)

	dst.id = src.id
	dst.scope = src.scope
	dst.loaded = false
	dst.info = src.info
	dst.instance = nil

	return dst
}

func (inst *innerComponentHolder) tryClone() *innerComponentHolder {
	if inst.scope == ScopePrototype {
		return inst.innerDoRealClone()
	}
	return inst
}

func (inst *innerComponentHolder) tryLoad() error {

	if inst.loaded {
		return nil
	}

	loader := inst.loader
	if loader == nil {
		loader = inst.context.manager.NewLoader()
		inst.loader = loader
	}

	com := loader.GetComponentByID(inst.id)
	err := loader.Load()
	if err != nil {
		return err
	}

	inst.instance = com
	inst.loaded = true
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentIndexer struct {
	context *innerComponentContext
	table   map[Selector]*innerComponentHolderChain
}

func (inst *innerComponentIndexer) init(ctx *innerComponentContext) {
	inst.context = ctx
	inst.table = make(map[Selector]*innerComponentHolderChain)
}

func (inst *innerComponentIndexer) putWithSelector(sel Selector, h *innerComponentHolder) {

	if sel == "" || h == nil {
		return
	}

	table := inst.table
	older := table[sel]

	node := new(innerComponentHolderChain)
	node.next = older
	node.sel = sel
	node.holder = h

	name := sel.String()
	if strings.HasPrefix(name, "#") {
		if older != nil {
			err := fmt.Errorf("innerComponentIndexer: selectors[%s] are duplicate", name)
			panic(err)
		}
	}

	table[sel] = node
}

func (inst *innerComponentIndexer) put(h *innerComponentHolder) {

	if h == nil {
		return
	}

	info := h.info
	if info == nil {
		return
	}

	id := info.ID
	inst.putWithSelector(id.Selector(), h)

	for _, key := range info.Aliases {
		inst.putWithSelector(key.Selector(), h)
	}

	for _, key := range info.Classes {
		inst.putWithSelector(key.Selector(), h)
	}

}

func (inst *innerComponentIndexer) selectList(sel Selector) []*innerComponentHolder {

	src := inst.selectChain(sel)
	dst := make([]*innerComponentHolder, 0)
	tmp := make(map[ID]*innerComponentHolder) // 用于防止si循环

	for p := src; p != nil; p = p.next {
		holder := p.holder
		if holder == nil {
			continue
		}
		id := holder.id
		older := tmp[id]
		if older != nil {
			break
		}
		tmp[id] = holder
		dst = append(dst, holder)
	}

	return dst
}

func (inst *innerComponentIndexer) selectOne(sel Selector) *innerComponentHolder {
	ch := inst.selectChain(sel)
	if ch == nil {
		return nil
	}
	return ch.holder
}

func (inst *innerComponentIndexer) selectChain(sel Selector) *innerComponentHolderChain {
	return inst.table[sel]
}

////////////////////////////////////////////////////////////////////////////////

type innerComponentContext struct {
	environment properties.Table
	props       properties.Table
	registry    Registry

	cache   *innerComponentCache
	indexer *innerComponentIndexer
	manager *innerComponentManager
}

// GetEnvironment implements Context.
func (inst *innerComponentContext) GetEnvironment() properties.Table {
	return inst.environment
}

// GetComponents implements Context.
func (inst *innerComponentContext) GetComponents() ComponentManager {
	return inst.manager
}

// GetProperties implements Context.
func (inst *innerComponentContext) GetProperties() properties.Table {
	return inst.props
}

func (inst *innerComponentContext) init() Context {

	inst.environment = properties.NewTable()
	inst.props = properties.NewTable()
	inst.registry = NewRegistry()
	inst.cache = new(innerComponentCache)
	inst.indexer = new(innerComponentIndexer)
	inst.manager = new(innerComponentManager)

	inst.cache.init(inst)
	inst.indexer.init(inst)
	inst.manager.init(inst)

	return inst
}

////////////////////////////////////////////////////////////////////////////////
