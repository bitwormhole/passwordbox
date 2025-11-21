package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
)

type PasswordComputer struct {

	// tmp
	ref   entity.PasswordBlock
	block entity.PasswordBlock

	// output
	plain dxo.PlainPassword
}

func (inst *PasswordComputer) Init(it *entity.PasswordBlock) error {

	// inst.entity = *it

	return nil
}

func (inst *PasswordComputer) Compute() error {

	return nil
}

func (inst *PasswordComputer) Result() dxo.PlainPassword {

	str := "todo ..."
	return dxo.PlainPassword(str)
}
