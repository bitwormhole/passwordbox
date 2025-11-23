package components

// Loader 表示组件的加载器
type Loader interface {
	GetComponentByID(id ID) any

	ListComponentsByClass(cls Class) []any
}

type Loading struct {
	Context   Context
	Loader    Loader
	Component any
	ID        ID
}

// 组件的工厂函数
type OnNewFunc func() any

// 组件的加载函数
type OnLoadFunc func(l *Loading)

// 组件的注册函数
type OnRegisterFunc func() *ComponentRegistration
