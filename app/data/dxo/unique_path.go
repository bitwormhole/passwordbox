package dxo

// UniquePath 表示密码的唯一（存储）路径
type UniquePath string

func (up UniquePath) String() string {
	return string(up)
}
