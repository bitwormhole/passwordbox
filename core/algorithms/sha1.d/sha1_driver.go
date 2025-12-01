package sha1d

import (
	"crypto/sha1"
	"fmt"

	"github.com/bitwormhole/passwordbox/core/algorithms"
)

////////////////////////////////////////////////////////////////////////////////

type Driver struct {
}

func (inst *Driver) AlgorithmName() algorithms.AlgorithmName {
	return algorithms.AlgorithmSHA1
}

func (inst *Driver) AlgorithmType() algorithms.AlgorithmType {
	return algorithms.TypeHash
}

func (inst *Driver) New() algorithms.Hash {
	return new(innerSHA1)
}

func (inst *Driver) Impl() algorithms.HashDriver {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerSHA1 struct {
	dr *Driver
}

func (i *innerSHA1) GetDriver() algorithms.HashDriver {
	return i.dr
}

func (i *innerSHA1) Sum(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}

func (i *innerSHA1) SumContext(ctx *algorithms.HashContext) error {

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
