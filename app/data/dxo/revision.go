package dxo

import "strconv"

type Revision int

func (r Revision) Int() int {
	return int(r)
}

func (r Revision) String() string {
	n := int(r)
	return strconv.Itoa(n)
}
