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
}

func (inst *PasswordServiceImpl) _impl() passwords.Service {
	return inst
}

func (inst *PasswordServiceImpl) Find(ctx context.Context, id dxo.PasswordID) (*dto.Password, error) {
	return nil, fmt.Errorf("no impl")
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
