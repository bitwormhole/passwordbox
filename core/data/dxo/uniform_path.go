package dxo

// UniformPath 表示资源的统一(存储）路径
type UniformPath string

func (up UniformPath) String() string {
	return string(up)
}
