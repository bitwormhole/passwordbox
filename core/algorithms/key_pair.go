package algorithms

type PublicKeyDriver interface {
	Driver

	GetPublicKeyLoader() PublicKeyLoader

	GetPrivateKeyLoader() PrivateKeyLoader

	GetPrivateKeyGenerator() PrivateKeyGenerator
}

type PublicKeyLoader interface {
	Load(p PEM) (PublicKey, error)
}

type PrivateKeyLoader interface {
	Load(p PEM) (PrivateKey, error)
}

type PrivateKeyGenerator interface {
	Generate() (PrivateKey, error)
}

////////////////////////////////////////////////////////////////////////////////

type PublicKey interface {
	Export() (PEM, error)

	GetDriver() PublicKeyDriver
}

type PrivateKey interface {
	Export() (PEM, error)

	GetDriver() PublicKeyDriver

	GetPublicKey() PublicKey
}
