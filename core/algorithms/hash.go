package algorithms

type HashDriver interface {
	Driver

	New() Hash
}

type Hash interface {
	GetDriver() HashDriver

	Sum(data []byte) []byte

	SumContext(ctx *HashContext) error
}

type HashContext struct {
	Algorithm AlgorithmName

	Driver HashDriver

	Data []byte

	Sum []byte

	Length uint
}
