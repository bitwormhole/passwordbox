package dxo

import (
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type PemFile struct {
	blocks []*pem.Block
}

func (inst *PemFile) Text() PemFileText {
	txt, err := inst.Format()
	if err != nil {
		vlog.Warn("error: PemFile.Text(): %v", err.Error())
	}
	return txt
}

func (inst *PemFile) Format() (PemFileText, error) {
	all := inst.blocks
	out := new(strings.Builder)
	for _, b := range all {
		err := pem.Encode(out, b)
		if err != nil {
			return "", err
		}
	}
	str := out.String()
	return PemFileText(str), nil
}

func (inst *PemFile) Add(b *pem.Block) {
	if b == nil {
		return
	}
	list := inst.blocks
	list = append(list, b)
	inst.blocks = list
}

func (inst *PemFile) FindBlock(t string) (*pem.Block, error) {
	s1 := strings.ToLower(t)
	all := inst.blocks
	for _, it := range all {
		s2 := strings.ToLower(it.Type)
		if s1 == s2 {
			return it, nil
		}
	}
	return nil, fmt.Errorf("PemFile.FindBlock(): no block with type of [%v]", s1)
}

func (inst *PemFile) Blocks() []*pem.Block {
	return inst.blocks
}

////////////////////////////////////////////////////////////////////////////////

type PemFileText string

func (txt PemFileText) String() string {
	return string(txt)
}

func (txt PemFileText) Parse(dst *PemFile) error {

	if dst == nil {
		return fmt.Errorf("PemFileText.Parse(): param 'dst' is nil")
	}

	parser := new(innerPemFileParser)
	f1, err := parser.parse([]byte(txt))

	if err != nil {
		return err
	}

	*dst = *f1
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type innerPemFileParser struct{}

func (inst *innerPemFileParser) parse(data []byte) (*PemFile, error) {

	list, err := inst.parseBlockList(data)
	if err != nil {
		return nil, err
	}

	pf := new(PemFile)
	pf.blocks = list
	return pf, nil
}

func (inst *innerPemFileParser) parseBlockList(raw []byte) ([]*pem.Block, error) {
	data := raw
	list := make([]*pem.Block, 0)
	for {
		b, rest := pem.Decode(data)
		if b != nil {
			list = append(list, b)
		} else {
			break
		}
		len2 := len(rest)
		if len2 < 1 {
			break
		}
		data = rest
	}
	return list, nil
}

////////////////////////////////////////////////////////////////////////////////
