package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {
	Find(db *gorm.DB, id dxo.PasswordID) (*entity.Password, error)

	Insert(db *gorm.DB, it *entity.Password) (*entity.Password, error)

	Update(db *gorm.DB, id dxo.PasswordID, fn func(it *entity.Password) error) (*entity.Password, error)

	Remove(db *gorm.DB, id dxo.PasswordID) error
}
