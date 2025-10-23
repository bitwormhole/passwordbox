package users

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"

	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.UserID) (*entity.User, error)

	// edit

	Insert(db *gorm.DB, item *entity.User) (*entity.User, error)

	Update(db *gorm.DB, id dxo.UserID, fn func(inner *entity.User) error) (*entity.User, error)

	Remove(db *gorm.DB, id dxo.UserID) error
}
