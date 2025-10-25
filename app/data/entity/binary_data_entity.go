package entity

import "github.com/bitwormhole/passwordbox/app/data/dxo"

////////////////////////////////////////////////////////////////////////////////

type BinaryData struct {

	// id

	ID dxo.BinaryDataID

	Base

	Type dxo.BinaryDataType

	Length int

	Sum dxo.BinaryDataSum

	ContentType dxo.MIMEType

	Content []byte
}

////////////////////////////////////////////////////////////////////////////////
// EOF
