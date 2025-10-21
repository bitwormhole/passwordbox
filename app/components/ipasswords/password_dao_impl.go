package ipasswords

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/database"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"gorm.io/gorm"
)

type PasswordDaoImpl struct {

	//starter:component

	_as func(passwords.DAO) //starter:as("#")

	Agent database.Agent //starter:inject("#")
}

func (inst *PasswordDaoImpl) _impl() passwords.DAO {
	return inst
}

func (inst *PasswordDaoImpl) Find(db *gorm.DB, id dxo.PasswordID) (*entity.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordDaoImpl) Insert(db *gorm.DB, it *entity.Password) (*entity.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordDaoImpl) Update(db *gorm.DB, id dxo.PasswordID, fn func(it *entity.Password) error) (*entity.Password, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordDaoImpl) Remove(db *gorm.DB, id dxo.PasswordID) error {

	return fmt.Errorf("no impl")
}
