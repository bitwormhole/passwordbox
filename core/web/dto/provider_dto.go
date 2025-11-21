package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

// Provider 表示一个在线服务 (passwordbox-hub-service) 提供商
type Provider struct {
	ID dxo.ProviderID `json:"id"`

	Base

	Domain dxo.DomainName `json:"domain"`

	Icon        dxo.URL `json:"icon"`
	WebSite     dxo.URL `json:"web"`
	Label       string  `json:"label"`
	Description string  `json:"description"`
}
