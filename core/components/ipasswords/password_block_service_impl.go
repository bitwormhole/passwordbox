package ipasswords

import (
	"context"
	"fmt"
	"sync"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type PasswordBlockServiceImpl struct {

	//starter:component

	_as func(passwords.BlockService) //starter:as("#")

	DaoChains passwords.ChainDAO //starter:inject("#")
	DaoBlocks passwords.BlockDAO //starter:inject("#")

	mutex sync.Mutex
}

func (inst *PasswordBlockServiceImpl) _impl() passwords.BlockService {
	return inst
}

func (inst *PasswordBlockServiceImpl) Find(ctx context.Context, id dxo.PasswordBlockID) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordBlockServiceImpl) Insert(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordBlockServiceImpl) CreateNewRevision(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	item.Revision = 0 // todo ...

	return inst.Insert(ctx, item)
}
