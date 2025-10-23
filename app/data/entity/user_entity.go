package entity

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type User struct {
	ID dxo.UserID

	Base

	Avatar      dxo.URL
	Email       dxo.EmailAddress `gorm:"unique"`
	ProviderDN  dxo.DomainName   // 服务提供商的域名
	DisplayName string

	KeyPair dxo.KeyPairID // 该用户的密钥对 ID
}
