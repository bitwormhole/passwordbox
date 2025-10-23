package users

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type Service interface {

	// query

	Find(ctx context.Context, id dxo.UserID) (*dto.User, error)

	// edit

	Insert(ctx context.Context, item *dto.User) (*dto.User, error)

	Update(ctx context.Context, id dxo.UserID, item *dto.User) (*dto.User, error)

	Remove(ctx context.Context, id dxo.UserID) error
}
