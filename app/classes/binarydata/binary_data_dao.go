package binarydata

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"gorm.io/gorm"
)

type DAO interface {

	// find

	Find(db *gorm.DB, id dxo.BinaryDataID) (*entity.BinaryData, error)

	// edit

	Insert(db *gorm.DB, item *entity.BinaryData) (*entity.BinaryData, error)

	Update(db *gorm.DB, id dxo.BinaryDataID, fn func(*entity.BinaryData) error) (*entity.BinaryData, error)

	Remove(db *gorm.DB, id dxo.BinaryDataID) error
}
