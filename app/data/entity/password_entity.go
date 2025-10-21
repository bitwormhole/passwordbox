package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
)

type Password struct {
	ID dxo.PasswordID

	Base

	Email    string
	Domain1  string
	Domain2  string
	UserName string
}
