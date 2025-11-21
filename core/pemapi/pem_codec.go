package pemapi

import (
	"bytes"
	"encoding/pem"
)

type CODEC struct {
}

func (inst *CODEC) decodePemDoc(raw []byte) ([]*pem.Block, error) {

	list := make([]*pem.Block, 0)
	data2do := raw

	for {
		block, rest := pem.Decode(data2do)
		if block == nil {
			break
		}
		list = append(list, block)
		data2do = rest
	}

	return list, nil
}

func (inst *CODEC) encodePemDoc(doc []*pem.Block) ([]byte, error) {
	out := new(bytes.Buffer)
	for _, block := range doc {
		err := pem.Encode(out, block)
		if err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}
