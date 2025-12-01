package backend

import "github.com/bitwormhole/passwordbox/core/pemapi"

type Router = pemapi.Router
type RouterConfig = pemapi.RouterConfig

// type RouterLoader = pemapi.RouterLoader

////////////////////////////////////////////////////////////////////////////////

// type innerRouterHolder struct {
// 	router *Router

// 	// config *RouterConfig
// 	// loader RouterLoader
// }

// func (inst *innerRouterHolder) getRouter() (*Router, error) {
// 	r := inst.router
// 	if r == nil {
// 		r2, err := inst.loadRouter()
// 		if err != nil {
// 			return nil, err
// 		}
// 		r = r2
// 		inst.router = r2
// 	}
// 	return r, nil
// }

// func (inst *innerRouterHolder) loadRouter() (*Router, error) {
// 	cfg := inst.getConfig()
// 	loader := cfg.Loader
// 	return loader.Load(cfg)
// }

// func (inst *innerRouterHolder) getConfig() *RouterConfig {

// 	config := new(RouterConfig)
// 	config.Handlers = pemapi.NewHandlerRegistry()
// 	config.Loader = pemapi.NewRouterLoader()

// 	// register handlers

// 	return config
// }

////////////////////////////////////////////////////////////////////////////////

// var theDefaultRouterHolder innerRouterHolder

// func GetDefaultRouter() (*Router, error) {
// 	h := &theDefaultRouterHolder
// 	return h.getRouter()
// }

////////////////////////////////////////////////////////////////////////////////
