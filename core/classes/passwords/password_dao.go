package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"gorm.io/gorm"
)

type BlockDAO interface {
	Find(db *gorm.DB, id dxo.PasswordBlockID) (*entity.PasswordBlock, error)

	FindByFingerPrint(db *gorm.DB, fp dxo.BlockFingerPrint) (*entity.PasswordBlock, error)

	Insert(db *gorm.DB, it *entity.PasswordBlock) (*entity.PasswordBlock, error)

	Update(db *gorm.DB, id dxo.PasswordBlockID, fn func(it *entity.PasswordBlock) error) (*entity.PasswordBlock, error)

	Remove(db *gorm.DB, id dxo.PasswordBlockID) error
}

type ChainDAO interface {
	Find(db *gorm.DB, id dxo.PasswordChainID) (*entity.PasswordChain, error)

	Insert(db *gorm.DB, it *entity.PasswordChain) (*entity.PasswordChain, error)

	Update(db *gorm.DB, id dxo.PasswordChainID, fn func(it *entity.PasswordChain) error) (*entity.PasswordChain, error)

	Remove(db *gorm.DB, id dxo.PasswordChainID) error
}
