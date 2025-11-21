package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type Example struct {
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
