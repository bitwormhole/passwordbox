package pemapi

import "strings"

type BlockType string

const (
	BlockTypeRequest     BlockType = "PEMAPI REQUEST"
	BlockTypeResponse    BlockType = "PEMAPI RESPONSE"
	BlockTypeEnvironment BlockType = "ENVIRONMENT"
	BlockTypeProperties  BlockType = "PROPERTIES"
)

func (bt BlockType) String() string {
	return string(bt)
}

func (bt BlockType) Equal(t2 BlockType) bool {
	t1 := bt.Normalize()
	t2 = t2.Normalize()
	return t1 == t2
}

func (bt BlockType) EqualString(t2 string) bool {
	return bt.Equal(BlockType(t2))
}

func (bt BlockType) Normalize() BlockType {
	str := bt.String()
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return BlockType(str)
}
