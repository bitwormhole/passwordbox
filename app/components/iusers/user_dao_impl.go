package iusers

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/users"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"gorm.io/gorm"
)

type UserDaoImpl struct {

	//starter:component

	_as func(users.DAO) //starter:as("#")

}

func (inst *UserDaoImpl) _impl() users.DAO {
	return inst
}

func (inst *UserDaoImpl) Find(db *gorm.DB, id dxo.UserID) (*entity.User, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *UserDaoImpl) Insert(db *gorm.DB, item *entity.User) (*entity.User, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *UserDaoImpl) Update(db *gorm.DB, id dxo.UserID, fn func(inner *entity.User) error) (*entity.User, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *UserDaoImpl) Remove(db *gorm.DB, id dxo.UserID) error {
	return fmt.Errorf("no impl")
}
