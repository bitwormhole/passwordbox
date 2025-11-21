package services

type Service interface {
	GetContext(dst *Context) *Context
}
