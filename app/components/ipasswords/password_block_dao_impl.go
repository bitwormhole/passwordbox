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

type PasswordBlockDaoImpl struct {

	//starter:component

	_as func(passwords.BlockDAO) //starter:as("#")

	Agent       database.Agent     //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")
}

func (inst *PasswordBlockDaoImpl) _impl() passwords.BlockDAO {
	return inst
}

func (inst *PasswordBlockDaoImpl) makeNewUUID() lang.UUID {
	b := inst.UUIDService.Build()
	b.Class("entity.Password")
	return b.Generate()
}

func (inst *PasswordBlockDaoImpl) Find(db *gorm.DB, id dxo.PasswordBlockID) (*entity.PasswordBlock, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordBlockDaoImpl) FindByFingerPrint(db *gorm.DB, fp dxo.BlockFingerPrint) (*entity.PasswordBlock, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *PasswordBlockDaoImpl) Insert(db *gorm.DB, it *entity.PasswordBlock) (*entity.PasswordBlock, error) {

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

func (inst *PasswordBlockDaoImpl) Update(db *gorm.DB, id dxo.PasswordBlockID, fn func(it *entity.PasswordBlock) error) (*entity.PasswordBlock, error) {

	return nil, fmt.Errorf("no impl")

}

func (inst *PasswordBlockDaoImpl) Remove(db *gorm.DB, id dxo.PasswordBlockID) error {

	return fmt.Errorf("no impl")
}
