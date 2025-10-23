package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
)

type PlainTextComputer struct {
	entity entity.Password
}

func (inst *PlainTextComputer) Init(it *entity.Password) error {

	inst.entity = *it

	return nil
}

func (inst *PlainTextComputer) Compute() dxo.PlainPassword {

	txt := ""

	return dxo.PlainPassword(txt)
}
