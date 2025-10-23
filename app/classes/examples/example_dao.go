package examples

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"

	"gorm.io/gorm"
)

type DAO interface {

	// query

	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)

	// edit

	Insert(db *gorm.DB, item *entity.Example) (*entity.Example, error)

	Update(db *gorm.DB, id dxo.ExampleID, fn func(inner *entity.Example) error) (*entity.Example, error)

	Remove(db *gorm.DB, id dxo.ExampleID) error
}
