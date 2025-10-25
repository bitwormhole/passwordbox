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

type PasswordChainDaoImpl struct {

	//starter:component

	_as func(passwords.ChainDAO) //starter:as("#")

	Agent       database.Agent     //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *PasswordChainDaoImpl) _impl() passwords.ChainDAO {
	return inst
}

func (inst *PasswordChainDaoImpl) makeNewUUID() lang.UUID {
	b := inst.UUIDService.Build()
	b.Class("entity.PasswordRef")
	return b.Generate()
}

func (inst *PasswordChainDaoImpl) Find(db *gorm.DB, id dxo.PasswordID) (*entity.PasswordChain, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordChainDaoImpl) Insert(db *gorm.DB, it *entity.PasswordChain) (*entity.PasswordChain, error) {

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

func (inst *PasswordChainDaoImpl) Update(db *gorm.DB, id dxo.PasswordID, fn func(it *entity.PasswordChain) error) (*entity.PasswordChain, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordChainDaoImpl) Remove(db *gorm.DB, id dxo.PasswordID) error {

	return fmt.Errorf("no impl")
}
