package uniformpaths

import (
	"strings"

	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

////////////////////////////////////////////////////////////////////////////////

func ComputePathForPasswordBlock(o *entity.PasswordBlock) dxo.UniformPath {

	upmkr := new(innerUniformPathMaker)
	upmkr.addElement("blocks")

	if o != nil {
		fp := o.Self
		upmkr.addElement(fp.String())
	}

	return upmkr.create()
}

func ComputePathForPasswordRef(o *entity.PasswordChain) dxo.UniformPath {

	upmkr := new(innerUniformPathMaker)
	upmkr.addElement("refs")

	if o != nil {
		upmkr.addElement(o.Email.String())
		upmkr.addElement(o.Domain.String())
		upmkr.addElement(o.Username)
		upmkr.addElement(o.Scene)
	}

	return upmkr.create()
}

func ComputePathForPassword(o *dto.Password) dxo.UniformPath {

	upmkr := new(innerUniformPathMaker)
	upmkr.addElement("words")

	if o != nil {
		upmkr.addElement(o.Email.String())
		upmkr.addElement(o.Domain.String())
		upmkr.addElement(o.UserName)
		upmkr.addElement(o.Scene)
	}

	return upmkr.create()
}

////////////////////////////////////////////////////////////////////////////////

type innerUniformPathMaker struct {
	builder strings.Builder
}

func (inst *innerUniformPathMaker) create() dxo.UniformPath {
	builder := &inst.builder
	str := builder.String()
	return dxo.UniformPath(str)
}

func (inst *innerUniformPathMaker) addElement(el string) {

	const sep = '/'

	el = strings.ReplaceAll(el, string(sep), "-")
	el = strings.TrimSpace(el)

	if len(el) < 1 {
		el = "undefined"
	}

	builder := &inst.builder
	if builder.Len() > 0 {
		builder.WriteRune(sep)
	}
	builder.WriteString(el)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
