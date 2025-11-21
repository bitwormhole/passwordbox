package entity

import (
	"github.com/bitwormhole/passwordbox/app/data/dxo"
)

type KeyPair struct {
	ID dxo.KeyPairID

	Base

	Algorithm string

	Fingerprint dxo.PublicKeyFingerPrint // 公钥指纹

	Size int // the key size in bits

	RawPublicKey dxo.BinaryDataRef

	RawPrivateKey dxo.BinaryDataRef
}
