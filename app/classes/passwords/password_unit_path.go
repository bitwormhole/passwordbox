package passwords

import (
	"strings"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

func MakeUniquePathWithDTO(item *dto.Password) dxo.UniquePath {
	mk := new(innerUniquePathMaker)
	mk.initWithDTO(item)
	return mk.create()
}

func MakeUniquePathWithEntity(item *entity.Password) dxo.UniquePath {
	mk := new(innerUniquePathMaker)
	mk.initWithEntity(item)
	return mk.create()
}

////////////////////////////////////////////////////////////////////////////////

type innerUniquePathMaker struct {
	builder strings.Builder
}

func (inst *innerUniquePathMaker) create() dxo.UniquePath {
	builder := &inst.builder
	str := builder.String()
	return dxo.UniquePath(str)
}

func (inst *innerUniquePathMaker) addElement(el string) {

	const sep = '/'

	el = strings.ReplaceAll(el, string(sep), "-")
	el = strings.TrimSpace(el)

	if len(el) < 1 {
		el = "undefined"
	}

	builder := &inst.builder
	builder.WriteRune(sep)
	builder.WriteString(el)
}

func (inst *innerUniquePathMaker) init(it *dto.Password) {

	rev := it.Revision.String()

	inst.addElement(it.Email.String())
	inst.addElement(it.Domain1.String())
	inst.addElement(it.Domain2.String())
	inst.addElement(it.UserName)
	inst.addElement(it.Scene)
	inst.addElement(rev)
}

func (inst *innerUniquePathMaker) initWithDTO(it1 *dto.Password) {
	it2 := new(dto.Password)
	if it1 != nil {
		*it2 = *it1
	}
	inst.init(it2)
}

func (inst *innerUniquePathMaker) initWithEntity(it1 *entity.Password) {
	it2 := new(dto.Password)
	if it1 != nil {
		ConvertE2D(it1, it2)
	}
	inst.init(it2)
}
