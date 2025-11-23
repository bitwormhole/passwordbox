package dxo

type EmailAddress string

func (addr EmailAddress) String() string {
	return string(addr)
}
