package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
)

type DigestComputer struct {
	entity entity.Password
}

func (inst *DigestComputer) Init(it *entity.Password) error {

	inst.entity = *it

	return nil
}

func (inst *DigestComputer) Compute() dxo.Digest {

	return ""
}
