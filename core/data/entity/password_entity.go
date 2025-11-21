package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/starter-go/base/lang"
)

// 表示四个维度的密码坐标
type PasswordCoordinate struct {
	Email    dxo.EmailAddress
	Domain   dxo.DomainName
	Username string
	Scene    string
}

// 用于生成密码的数据块
type PasswordBlock struct {

	// ids

	ID dxo.PasswordBlockID

	Base

	PasswordCoordinate

	// refer

	Chain dxo.PasswordChainID // 所属密码链

	// parent

	Parent dxo.BlockFingerPrint // 上个版本的指纹

	Self dxo.BlockFingerPrint // 这个版本的指纹

	// params(this)

	Revision dxo.Revision
	Charset  string
	Length   int
	Salt     lang.Hex

	// content

	Secret dxo.BinaryDataRef // ( 生成模式：为空 | 存储模式：经过加密的用户密码 )

	Content dxo.BinaryDataRef // 用于计算摘要的内容
}

// 对 密码链的头部 引用
type PasswordChain struct {

	// ids

	ID dxo.PasswordChainID

	Base

	PasswordCoordinate

	// HEAD

	Head dxo.BlockFingerPrint
}
