package ipasswords

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
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
	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordServiceImpl) Update(ctx context.Context, id dxo.PasswordID, item *dto.Password) (*dto.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordServiceImpl) Remove(ctx context.Context, id dxo.PasswordID) error {

	return fmt.Errorf("no impl")
}
