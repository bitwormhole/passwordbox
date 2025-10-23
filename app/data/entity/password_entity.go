package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type Password struct {

	// ids

	ID     dxo.PasswordID
	Parent dxo.PasswordID // 上一个版本的 ID
	Head   dxo.PasswordID // 当前版本的 ID ( 仅仅当这个 entity 的 Revision==0 时有效 )
	Root   dxo.PasswordID // 版本 '0' 的 ID

	Base

	// params(base)

	Email    dxo.EmailAddress
	Domain1  dxo.DomainName
	Domain2  dxo.DomainName
	UserName string
	Scene    string
	Revision dxo.Revision

	// params(ext)

	Charset string
	Length  int
	Salt    lang.Base64
	Word    lang.Base64 // ( 生成模式：为空 | 存储模式：经过加密的用户密码 )

	// others

	Path dxo.UniquePath `gorm:"unique"` // 这个实体的唯一路径

	ParentDigest dxo.Digest // 上一个版本的摘要
	ThisDigest   dxo.Digest // 这个版本的摘要

	Latest      dxo.Revision    // 最新版本（ 仅当 entity 的 Revision==0 时有效 ）
	AccountType dxo.AccountType // 表示账号类型
}
