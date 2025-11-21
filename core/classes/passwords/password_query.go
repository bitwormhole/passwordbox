package passwords

import (
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/starter-go/rbac"
)

type Query struct {
	All bool

	Pagination rbac.Pagination

	Want *dto.Password
}
