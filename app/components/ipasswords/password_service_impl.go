package ipasswords

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type PasswordServiceImpl struct {

	//starter:component

	_as func(passwords.Service) //starter:as("#")

	Dao passwords.DAO //starter:inject("#")
}

func (inst *PasswordServiceImpl) _impl() passwords.Service {
	return inst
}

func (inst *PasswordServiceImpl) Find(ctx context.Context, id dxo.PasswordID) (*dto.Password, error) {

	dao := inst.Dao

	it1, err := dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	it2 := new(dto.Password)

	err = passwords.ConvertE2D(it1, it2)
	if err != nil {
		return nil, err
	}

	return it2, nil
}

func (inst *PasswordServiceImpl) Query(ctx context.Context, q *passwords.Query) ([]*dto.Password, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordServiceImpl) Insert(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	dao := inst.Dao
	it2 := new(entity.Password)
	it4 := new(dto.Password)

	err := passwords.ConvertD2E(item, it2)
	if err != nil {
		return nil, err
	}

	it3, err := dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	err = passwords.ConvertE2D(it3, it4)
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

func (inst *PasswordServiceImpl) InitNewAccount(ctx context.Context, item *dto.Password, at dxo.AccountType) (*dto.Password, error) {

	item.Revision = 0 // todo ...

	if at == dxo.AccountTypeMaster {

		item.Domain2 = ""
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

func (inst *PasswordServiceImpl) CreateNewRevision(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	item.Revision = 0 // todo ...

	return inst.Insert(ctx, item)
}
