package mo

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type BinaryData struct {

	// id

	ID dxo.BinaryDataID

	Base

	Sum      dxo.BinaryDataSum
	DataType dxo.BinaryDataType
	MimeType dxo.MIMEType

	Length int

	Content []byte
}
