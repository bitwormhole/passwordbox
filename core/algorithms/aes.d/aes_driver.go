package aesd

import "github.com/bitwormhole/passwordbox/core/algorithms"

type Driver struct {
}

// AlgorithmName implements algorithms.SecretKeyDriver.
func (inst *Driver) AlgorithmName() algorithms.AlgorithmName {
	panic("unimplemented")
}

// AlgorithmType implements algorithms.SecretKeyDriver.
func (inst *Driver) AlgorithmType() algorithms.AlgorithmType {
	panic("unimplemented")
}

func (inst *Driver) _impl() algorithms.SecretKeyDriver {
	return inst
}
