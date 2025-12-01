package sha256d

import (
	"crypto/sha256"
	"fmt"

	"github.com/bitwormhole/passwordbox/core/algorithms"
)

type Driver struct {
}

func (inst *Driver) AlgorithmName() algorithms.AlgorithmName {
	return algorithms.AlgorithmSHA256
}

func (inst *Driver) AlgorithmType() algorithms.AlgorithmType {
	return algorithms.TypeHash
}

func (inst *Driver) New() algorithms.Hash {
	return new(innerSHA256)
}

func (inst *Driver) Impl() algorithms.HashDriver {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerSHA256 struct {
	dr *Driver
}

func (i *innerSHA256) GetDriver() algorithms.HashDriver {
	return i.dr
}

func (i *innerSHA256) Sum(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func (i *innerSHA256) SumContext(ctx *algorithms.HashContext) error {

	if ctx == nil {
		return fmt.Errorf("param:context is nil")
	}

	data := ctx.Data
	total := len(data)

	ctx.Algorithm = i.dr.AlgorithmName()
	ctx.Length = uint(total)
	ctx.Sum = i.Sum(data)
	ctx.Driver = i.dr

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
