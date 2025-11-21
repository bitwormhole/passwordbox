package binarydata

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/mo"
)

type Service interface {

	// find

	Find(cc context.Context, id dxo.BinaryDataID) (*mo.BinaryData, error)

	// edit

	Insert(cc context.Context, item *mo.BinaryData) (*mo.BinaryData, error)
}
