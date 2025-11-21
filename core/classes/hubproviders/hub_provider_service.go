package hubproviders

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type Service interface {

	// query

	Find(ctx context.Context, id dxo.ProviderID) (*dto.Provider, error)

	ListAll(ctx context.Context) ([]*dto.Provider, error)

	// edit

	Insert(ctx context.Context, item *dto.Provider) (*dto.Provider, error)

	Update(ctx context.Context, id dxo.ProviderID, item *dto.Provider) (*dto.Provider, error)

	Remove(ctx context.Context, id dxo.ProviderID) error
}
