package algorithms

import "encoding/pem"

type PEM struct {
	Blocks []*pem.Block
}
