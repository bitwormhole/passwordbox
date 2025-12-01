package pemapi

import (
	"fmt"
	"strings"
	"sync"
)

////////////////////////////////////////////////////////////////////////////////

type innerRouterRaw struct {
	handlers []*HandlerRegistration
	filters  []*HandlerRegistration
}

func (inst *innerRouterRaw) init() {

	inst.filters = make([]*HandlerRegistration, 0)
	inst.handlers = make([]*HandlerRegistration, 0)
}

////////////////////////////////////////////////////////////////////////////////

type innerRouterCache struct {
	mutex sync.Mutex

	handlersWildcard []*HandlerHolder
	handlersAll      []*HandlerHolder
	handlersMap      map[HandlerKey]*HandlerHolder

	chain FilterChain
}

func (inst *innerRouterCache) load(r *Router) error {

	err := inst.innerLoadFilters(r)
	if err != nil {
		return err
	}

	err = inst.innerLoadHandlers(r)
	if err != nil {
		return err
	}

	return nil
}

func (inst *innerRouterCache) innerLoadFilters(r *Router) error {
	src := r.raw.filters
	chainBuilder := new(FilterChainBuilder)
	chainBuilder.Add(src...)
	inst.chain = chainBuilder.Build()
	return nil
}

func (inst *innerRouterCache) innerLoadHandlers(r *Router) error {

	table := make(map[HandlerKey]*HandlerHolder)
	wildcards := inst.handlersWildcard
	all := inst.handlersAll
	src := r.raw.handlers

	for _, hr := range src {
		holder := inst.innerMakeHandlerHolder(hr)
		key := holder.Key
		if key.IsWildCard() {
			wildcards = append(wildcards, holder)
		}
		all = append(all, holder)
		table[key] = holder
	}

	inst.handlersWildcard = wildcards
	inst.handlersAll = all
	inst.handlersMap = table
	return nil
}

func (inst *innerRouterCache) innerMakeHandlerHolder(hr *HandlerRegistration) *HandlerHolder {

	hkb := new(HandlerKeyBuilder)
	dst := new(HandlerHolder)

	hkb.Method = hr.Method
	hkb.Path = hr.Path

	dst.Handler = hr.Handler
	dst.Key = hkb.Key()

	return dst
}

func (inst *innerRouterCache) findHandler(key HandlerKey) *HandlerHolder {
	res := inst.innerFindHandlerInMap(key)
	if res != nil {
		return res
	}
	return inst.innerFindHandlerInWildcards(key)
}

func (inst *innerRouterCache) innerFindHandlerInMap(key HandlerKey) *HandlerHolder {

	mtx := &inst.mutex
	mtx.Lock()
	defer mtx.Unlock()

	return inst.handlersMap[key]
}

func (inst *innerRouterCache) innerFindHandlerInWildcards(key HandlerKey) *HandlerHolder {
	list := inst.handlersWildcard
	for _, hh := range list {
		if key.Match(hh.Key) {
			return hh
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type Router struct {
	raw    *innerRouterRaw
	cache  *innerRouterCache
	config *RouterConfig

	Controllers []Controller // inject
}

func (inst *Router) init(cfg *RouterConfig) {

	raw := new(innerRouterRaw)
	raw.init()

	inst.raw = raw
	inst.config = cfg
}

// func (inst *Router) ListRegistrations() []*HandlerRegistration {
// 	return nil
// }

func (inst *Router) Handle(c *Context) error {

	err := inst.innerParseRequest(c)
	if err != nil {
		return err
	}

	key := c.Key
	hh, err := inst.innerFindHandler(key)
	if err != nil {
		return err
	}

	cache := inst.innerGetCache()
	c.Handler = hh.Handler
	c.Filters = cache.chain

	return c.Filters.Pass(c)
}

func (inst *Router) innerParseRequest(c *Context) error {

	req := c.Request
	location := req.Location
	method := req.Method

	hkb := new(HandlerKeyBuilder)
	hkb.Method = method
	hkb.Path = location.Path
	key := hkb.Key()

	req.Protocol = location.Scheme
	req.User = location.User.Username()
	req.Host = location.Host
	req.PathWant = location.Path
	req.PathHave = ""

	c.Key = key

	return nil
}

func (inst *Router) innerGetHandler() Handler {
	return inst
}

func (inst *Router) innerGetCache() *innerRouterCache {
	cache := inst.cache
	if cache != nil {
		return cache
	}
	cache, err := inst.innerLoadCache()
	if err != nil {
		panic(err)
	}
	inst.cache = cache
	return cache
}

func (inst *Router) innerLoadCache() (*innerRouterCache, error) {
	cache := new(innerRouterCache)
	err := cache.load(inst)
	if err != nil {
		return nil, err
	}
	return cache, nil
}

// func (inst *Router) tryAdd() *innerRouterCache {
// 	return inst
// }

func (inst *Router) innerFindHandler(key HandlerKey) (*HandlerHolder, error) {
	cache := inst.innerGetCache()
	h := cache.findHandler(key)
	if h != nil {
		return h, nil
	}
	return nil, fmt.Errorf("pemapi.Router: no handler for api [%s]", key)
}

func (inst *Router) AddHandlerFn(method string, path string, fn HandlerFunc) {
	ada := new(HandlerAdapter)
	ada.Fn = fn
	inst.AddHandler(method, path, ada)
}

func (inst *Router) AddFilterFn(method string, path string, fn FilterFunc) {
	ada := new(FilterAdapter)
	ada.Fn = fn
	inst.AddFilter(method, path, ada)
}

func (inst *Router) AddHandler(method string, path string, h Handler) {

	r1 := new(HandlerRegistration)
	r1.Handler = h

	r1.Method = method
	r1.Path = path
	r1.Priority = 0
	r1.Enabled = true
	r1.Name = method + ":" + path

	raw := inst.raw
	raw.handlers = append(raw.handlers, r1)
}

func (inst *Router) AddFilter(method string, path string, f Filter) {

	r1 := new(HandlerRegistration)
	r1.Filter = f

	r1.Method = method
	r1.Path = path
	r1.Priority = 0
	r1.Enabled = true
	r1.Name = method + ":" + path

	raw := inst.raw
	raw.filters = append(raw.filters, r1)
}

////////////////////////////////////////////////////////////////////////////////

type HandlerKey string

func (k HandlerKey) String() string {
	return string(k)
}

// 判断是否为通配符
func (k HandlerKey) IsWildCard() bool {

	str := k.String()

	if strings.HasPrefix(str, ":") {
		return true
	}

	if strings.Contains(str, "/:") {
		return true
	}

	return strings.ContainsAny(str, "?*")
}

// 检查两个键是否匹配
func (k HandlerKey) Match(wildcard HandlerKey) bool {

	elist1 := k.Elements()
	elist2 := wildcard.Elements()

	if len(elist1) != len(elist2) {
		return false
	}

	for idx, e1 := range elist1 {
		e2 := elist2[idx]
		if e2.IsWildCard() {
			continue
		} else {
			if e1 != e2 {
				return false
			}
		}
	}

	return true
}

// 拆分成元素列表
func (k HandlerKey) Elements() []HandlerKey {

	const sep = "/"
	str := k.String()
	src := strings.Split(str, sep)
	dst := make([]HandlerKey, 0)

	for _, item := range src {
		item = strings.TrimSpace(item)
		if len(item) > 0 {
			dst = append(dst, HandlerKey(item))
		}
	}

	return dst
}

// 提取参数
func (k HandlerKey) ExtractParams(wildcard HandlerKey) map[string]string {

	dst := make(map[string]string)
	elist1 := k.Elements()
	elist2 := wildcard.Elements()

	if len(elist1) != len(elist2) {
		return dst
	}

	for idx, e1 := range elist1 {
		e2 := elist2[idx]
		if e2.IsWildCard() {
			name := e2.String()
			value := e1.String()
			if strings.HasPrefix(name, ":") {
				name = name[1:]
			}
			dst[name] = value
		}
	}

	return dst
}

////////////////////////////////////////////////////////////////////////////////

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
	// Handlers HandlerRegistry
	// Loader RouterLoader

	foo int
}

////////////////////////////////////////////////////////////////////////////////

// type RouterLoader interface {
// 	Load(config *RouterConfig) (*Router, error)
// }

// func NewRouterLoader() RouterLoader {
// 	loader := new(innerRouterLoader)
// 	return loader
// }

////////////////////////////////////////////////////////////////////////////////

// type innerRouterLoader struct{}

// // Load implements RouterLoader.
// func (inst *innerRouterLoader) Load(config *RouterConfig) (*Router, error) {

// 	rt := new(Router)
// 	rt.init(config)

// 	src := config.Handlers.ListAll()
// 	hkb := new(HandlerKeyBuilder)

// 	for _, item := range src {
// 		if !inst.innerIsHandlerReady(item) {
// 			continue
// 		}
// 		hkb.Path = item.Path
// 		hkb.Method = item.Method
// 		key := hkb.Key()

// 		rt.handlers[key] = item
// 	}

// 	return rt, nil
// }

// func (i *innerRouterLoader) innerIsHandlerReady(item *HandlerRegistration) bool {

// 	if item == nil {
// 		return false
// 	}

// 	if item.Path == "" {
// 		return false
// 	}
// 	if item.Handler == nil {
// 		return false
// 	}
// 	if !item.Enabled {
// 		return false
// 	}

// 	return true
// }

////////////////////////////////////////////////////////////////////////////////
