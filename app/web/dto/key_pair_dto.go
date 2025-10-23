package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type KeyPair struct {
	ID dxo.KeyPairID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
