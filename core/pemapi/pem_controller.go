package pemapi

type Controller interface {
	Route(router *Router) error
}
