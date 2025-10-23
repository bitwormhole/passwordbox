package examples

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type Service interface {

	// query

	Find(ctx context.Context, id dxo.ExampleID) (*dto.Example, error)

	// edit

	Insert(ctx context.Context, item *dto.Example) (*dto.Example, error)

	Update(ctx context.Context, id dxo.ExampleID, item *dto.Example) (*dto.Example, error)

	Remove(ctx context.Context, id dxo.ExampleID) error
}
