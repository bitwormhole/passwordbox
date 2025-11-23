package components

type ID string
type Class string
type Alias string
type Scope int

const (
	ScopeSingleton Scope = 1
	ScopePrototype Scope = 2
)

type ComponentManager interface {
	GetComponent(id ID) any

	TryGetComponent(id ID) (any, error)
}
