package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/starter-go/base/lang"
)

type KeyPair struct {
	ID dxo.KeyPairID

	Base

	Algorithm string

	FingerPrint lang.Hex // 公钥指纹

	KeySize int // key size in bits

	PublicKey  dxo.PemFileText
	PrivateKey dxo.PemFileText
}
