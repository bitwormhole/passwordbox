package ipasswords

import (
	"context"
	"fmt"
	"sync"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type PasswordChainServiceImpl struct {

	//starter:component

	_as func(passwords.ChainService) //starter:as("#")

	DaoChains passwords.ChainDAO //starter:inject("#")
	DaoBlocks passwords.BlockDAO //starter:inject("#")

	mutex sync.Mutex
}

func (inst *PasswordChainServiceImpl) _impl() passwords.ChainService {
	return inst
}

func (inst *PasswordChainServiceImpl) Find(ctx context.Context, id dxo.PasswordChainID) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordChainServiceImpl) Insert(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordChainServiceImpl) InitNewAccount(ctx context.Context, item *dto.Password, at dxo.AccountType) (*dto.Password, error) {

	item.Revision = 0 // todo ...

	if at == dxo.AccountTypeMaster {

		item.UserName = ""
		item.Scene = ""
		item.Revision = 0

		item.AccountType = at

	} else if at == dxo.AccountTypeSlave {

		item.Revision = 0
		item.AccountType = at

	} else {
		return nil, fmt.Errorf("PasswordServiceImpl.InitNewAccount: bad AccountType")
	}

	return inst.Insert(ctx, item)
}
