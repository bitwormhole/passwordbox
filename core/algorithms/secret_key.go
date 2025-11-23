package algorithms

type SecretKeyDriver interface {
	Driver
}

type SecretKeyLoader interface {
	Load(p PEM) (SecretKey, error)
}

type SecretKeyGenerator interface {
	Generate() SecretKey
}

type SecretKey interface {
	GetDriver() SecretKeyDriver

	Export() (PEM, error)
}
