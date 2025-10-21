package entity

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type Example struct {
	ID dxo.ExampleID

	Base

	Foo string
	Bar int
}
