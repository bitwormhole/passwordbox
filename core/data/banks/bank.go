package banks

import "github.com/bitwormhole/passwordbox/core/algorithms"

type BankName string

type Bank interface {
	GetName() BankName

	GetConfigurations() Configurations

	GetRefs() Refs

	GetObjects() Objects

	GetKeys() BankKeys

	Init(opt *BankInitOptions) error

	Delete() error

	Exists() bool
}

type BankInitOptions struct {
	Name BankName

	SecretKey algorithms.SecretKey
}

type BankKeys struct {
	PrivateKey algorithms.PrivateKey
	PublicKey  algorithms.PublicKey
	SecretKey  algorithms.SecretKey
}
