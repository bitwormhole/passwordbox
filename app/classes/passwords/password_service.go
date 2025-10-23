package passwords

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type Service interface {
	Find(ctx context.Context, id dxo.PasswordID) (*dto.Password, error)

	Query(ctx context.Context, q *Query) ([]*dto.Password, error)

	Insert(ctx context.Context, item *dto.Password) (*dto.Password, error)

	Update(ctx context.Context, id dxo.PasswordID, item *dto.Password) (*dto.Password, error)

	Remove(ctx context.Context, id dxo.PasswordID) error

	// 初始化一个新的账号
	InitNewAccount(ctx context.Context, item *dto.Password, at dxo.AccountType) (*dto.Password, error)

	// 创建一个新的密码版本
	CreateNewRevision(ctx context.Context, item *dto.Password) (*dto.Password, error)
}
