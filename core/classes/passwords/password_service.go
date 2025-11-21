package passwords

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"

	"github.com/bitwormhole/passwordbox/app/web/dto"
)

// the main-service
type Service interface {
	Find(ctx context.Context, id dxo.PasswordID) (*dto.Password, error)

	Query(ctx context.Context, q *Query) ([]*dto.Password, error)

	Insert(ctx context.Context, item *dto.Password) (*dto.Password, error)

	Update(ctx context.Context, id dxo.PasswordID, item *dto.Password) (*dto.Password, error)

	Remove(ctx context.Context, id dxo.PasswordID) error
}

// the chain-service
type ChainService interface {

	// 初始化一个新的账号
	InitNewAccount(ctx context.Context, item *dto.Password, at dxo.AccountType) (*dto.Password, error)

	Find(ctx context.Context, id dxo.PasswordChainID) (*dto.Password, error)

	Insert(ctx context.Context, item *dto.Password) (*dto.Password, error)
}

// the block-service
type BlockService interface {

	// 创建一个新的密码版本
	CreateNewRevision(ctx context.Context, item *dto.Password) (*dto.Password, error)

	Find(ctx context.Context, id dxo.PasswordBlockID) (*dto.Password, error)

	Insert(ctx context.Context, item *dto.Password) (*dto.Password, error)
}

// the fast-service
type FastService interface {

	// 根据传入的参数， 快速生成（或查询）一个密码
	MakeFastGen(ctx context.Context, item *dto.Password) (*dto.Password, error)
}
