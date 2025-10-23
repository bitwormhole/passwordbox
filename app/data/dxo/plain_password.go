package dxo

type PlainPassword []rune

func (pp PlainPassword) String() string {
	return string(pp)
}
