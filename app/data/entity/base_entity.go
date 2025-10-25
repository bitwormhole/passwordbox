package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/starter-go/security-gorm/rbacdb"
)

type Base struct {
	rbacdb.BaseEntity

	Path dxo.UniformPath `gorm:"unique"` // 这个实体的唯一路径

}
