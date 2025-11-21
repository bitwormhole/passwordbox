package ipasswords

import (
	"context"
	"fmt"
	"sync"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type PasswordServiceImpl struct {

	//starter:component

	_as func(passwords.Service) //starter:as("#")

	DaoChains passwords.ChainDAO //starter:inject("#")
	DaoBlocks passwords.BlockDAO //starter:inject("#")

	mutex sync.Mutex
}

func (inst *PasswordServiceImpl) _impl() passwords.Service {
	return inst
}

func (inst *PasswordServiceImpl) Find(ctx context.Context, id dxo.PasswordID) (*dto.Password, error) {

	dao1 := inst.DaoChains
	dao2 := inst.DaoBlocks

	it1a, err := dao1.Find(nil, id)
	if err != nil {
		return nil, err
	}

	head := it1a.Head
	it1b, err := dao2.FindByFingerPrint(nil, head)
	if err != nil {
		return nil, err
	}

	it2 := new(dto.Password)
	err = passwords.ConvertE2D(it1a, it1b, it2)
	if err != nil {
		return nil, err
	}

	return it2, nil
}

func (inst *PasswordServiceImpl) Query(ctx context.Context, q *passwords.Query) ([]*dto.Password, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordServiceImpl) Insert(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	dao1 := inst.DaoChains
	dao2 := inst.DaoBlocks

	it2a := new(entity.PasswordChain)
	it2b := new(entity.PasswordBlock)
	it4 := new(dto.Password)

	err := passwords.ConvertD2E(item, it2a, it2b)
	if err != nil {
		return nil, err
	}

	it3a, err := dao1.Insert(nil, it2a)
	if err != nil {
		return nil, err
	}
	it3b, err := dao2.Insert(nil, it2b)
	if err != nil {
		return nil, err
	}

	err = passwords.ConvertE2D(it3a, it3b, it4)
	if err != nil {
		return nil, err
	}

	return it4, nil
}

func (inst *PasswordServiceImpl) Update(ctx context.Context, id dxo.PasswordID, item *dto.Password) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordServiceImpl) Remove(ctx context.Context, id dxo.PasswordID) error {

	return fmt.Errorf("no impl")
}
