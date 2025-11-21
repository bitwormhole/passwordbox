package dxo

type DomainName string

func (dn DomainName) String() string {
	return string(dn)
}
