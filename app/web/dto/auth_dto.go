package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type Auth struct {
	ID dxo.AuthID `json:"id"`

	Base

	Mechanism     string `json:"mechanism"`
	Step          int    `json:"step"`
	UserKeyID     string `json:"user_key_id"`
	UserKeySecret string `json:"user_key_secret"`

	UserEmail  dxo.EmailAddress `json:"user_email"`
	ProviderDN dxo.DomainName   `json:"provider_domain"`
}
