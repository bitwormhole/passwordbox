package iusers

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/users"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type UserServiceImpl struct {

	//starter:component

	_as func(users.Service) //starter:as("#")

}

func (inst *UserServiceImpl) _impl() users.Service {
	return inst
}

func (inst *UserServiceImpl) Find(ctx context.Context, id dxo.UserID) (*dto.User, error) {
	return nil, fmt.Errorf("no impl")

}

func (inst *UserServiceImpl) Insert(ctx context.Context, item *dto.User) (*dto.User, error) {
	return nil, fmt.Errorf("no impl")

}

func (inst *UserServiceImpl) Update(ctx context.Context, id dxo.UserID, item *dto.User) (*dto.User, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *UserServiceImpl) Remove(ctx context.Context, id dxo.UserID) error {

	return fmt.Errorf("no impl")
}
