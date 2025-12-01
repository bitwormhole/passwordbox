package rsad

import "github.com/bitwormhole/passwordbox/core/algorithms"

type Driver struct {
}

// AlgorithmName implements algorithms.PublicKeyDriver.
func (inst *Driver) AlgorithmName() algorithms.AlgorithmName {
	panic("unimplemented")
}

// AlgorithmType implements algorithms.PublicKeyDriver.
func (inst *Driver) AlgorithmType() algorithms.AlgorithmType {
	panic("unimplemented")
}

// GetPrivateKeyGenerator implements algorithms.PublicKeyDriver.
func (inst *Driver) GetPrivateKeyGenerator() algorithms.PrivateKeyGenerator {
	panic("unimplemented")
}

// GetPrivateKeyLoader implements algorithms.PublicKeyDriver.
func (inst *Driver) GetPrivateKeyLoader() algorithms.PrivateKeyLoader {
	panic("unimplemented")
}

// GetPublicKeyLoader implements algorithms.PublicKeyDriver.
func (inst *Driver) GetPublicKeyLoader() algorithms.PublicKeyLoader {
	panic("unimplemented")
}

func (inst *Driver) _impl() algorithms.PublicKeyDriver {
	return inst
}
