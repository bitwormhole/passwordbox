package digests

import (
	"crypto/sha256"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
)

func Compute(data []byte) dxo.Digest {
	sum := sha256.Sum256(data)
	return sum
}
