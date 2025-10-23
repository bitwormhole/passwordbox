package ipasswords

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/data/database"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type PasswordDaoImpl struct {

	//starter:component

	_as func(passwords.DAO) //starter:as("#")

	Agent       database.Agent     //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *PasswordDaoImpl) _impl() passwords.DAO {
	return inst
}

func (inst *PasswordDaoImpl) makeNewUUID() lang.UUID {
	b := inst.UUIDService.Build()
	b.Class("entity.Password")
	return b.Generate()
}

func (inst *PasswordDaoImpl) Find(db *gorm.DB, id dxo.PasswordID) (*entity.Password, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordDaoImpl) Insert(db *gorm.DB, it *entity.Password) (*entity.Password, error) {

	db = inst.Agent.DB(db)
	uuid := inst.makeNewUUID()

	it.ID = 0
	it.UUID = uuid

	res := db.Create(it)
	err := res.Error
	if err != nil {
		return nil, err
	}
	return it, nil
}

func (inst *PasswordDaoImpl) Update(db *gorm.DB, id dxo.PasswordID, fn func(it *entity.Password) error) (*entity.Password, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordDaoImpl) Remove(db *gorm.DB, id dxo.PasswordID) error {

	return fmt.Errorf("no impl")
}
