package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type Password struct {
	ID dxo.PasswordID `json:"id"`

	Base

	// pass-path: email/domain1/domain2/username/scene/revision

	Domain1  string `json:"domain1"`  // 提供托管服务的域名
	Domain2  string `json:"domain2"`  // 被托管的目标域名
	Email    string `json:"email"`    // 用户的邮箱地址
	UserName string `json:"username"` // 用户名
	Scene    string `json:"scene"`    // 该密码的使用场景
	CharSet  string `json:"charset"`  // 构成密码的字符集

	Verification string `json:"verification"` // 验证码 ( 可以是动态的，也可以是静态的 )
	PlainWord    string `json:"word"`         // 密码的明文 （仅在存储模式下用于输入）

	Revision int `json:"revision"` // 密码的版本
	Length   int `json:"length"`   // 密码的长度
}
