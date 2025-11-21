package examples

import (
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/starter-go/rbac"
)

type Query struct {
	All bool

	Pagination rbac.Pagination

	Want *entity.Example
}
