package banks

import "github.com/bitwormhole/passwordbox/core/data/dxo"

type Store interface {
	GetBank(name BankName) Bank

	GetRootBank() Bank

	GetUserBank(addr dxo.EmailAddress) (Bank, error)
}

func GetStore() Store {
	return nil
}
