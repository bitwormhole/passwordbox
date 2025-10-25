package ipasswords

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/classes/services"
	"github.com/bitwormhole/passwordbox/app/classes/uniformpaths"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/data/mo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/vlog"
)

type innerFastPasswordGenerator struct {

	// contexts

	cc context.Context
	sc *services.Context

	out strings.Builder

	tmpRef   entity.PasswordChain
	tmpBlock entity.PasswordBlock
}

func (inst *innerFastPasswordGenerator) Init(it1 *entity.PasswordChain, it2 *entity.PasswordBlock) {
	inst.tmpRef = *it1
	inst.tmpBlock = *it2
}

func (inst *innerFastPasswordGenerator) InitWithDTO(it *dto.Password) {

	it1 := new(entity.PasswordChain)
	it2 := new(entity.PasswordBlock)

	passwords.ConvertD2E(it, it1, it2)
	inst.Init(it1, it2)
}

func (inst *innerFastPasswordGenerator) InitWithEntity(it1 *entity.PasswordChain, it2 *entity.PasswordBlock) {
	inst.Init(it1, it2)
}

func (inst *innerFastPasswordGenerator) Generate() (dxo.PlainPassword, error) {

	str := ""
	err := inst.innerGenerate()
	if err != nil {
		return dxo.PlainPassword(str), err
	}

	str = inst.out.String()
	return dxo.PlainPassword(str), nil
}

func (inst *innerFastPasswordGenerator) innerGenerate() error {

	// tmp1 := &inst.tmpRef
	// tmp2 := &inst.tmpBlock

	steps := make([]func(stepIndex int) error, 0)

	steps = append(steps, inst.doStepMakeContent)
	steps = append(steps, inst.doStepMakeDigest)
	steps = append(steps, inst.doStepMakePath)
	steps = append(steps, inst.doStepMakePassword)

	for i, fn := range steps {
		err := fn(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *innerFastPasswordGenerator) doStep(step int) error {
	return fmt.Errorf("innerFastPasswordGenerator: no impl")
}

func (inst *innerFastPasswordGenerator) makeContent(builder *strings.Builder, name string, value string) {
	builder.WriteString(name)
	builder.WriteRune(':')
	builder.WriteString(value)
	builder.WriteRune('\n')
}

func (inst *innerFastPasswordGenerator) doStepMakeContent(step int) error {

	block := &inst.tmpBlock
	ref := &inst.tmpRef
	builder := new(strings.Builder)

	length := strconv.Itoa(block.Length)
	charset := block.Charset
	domain := ref.Domain
	email := ref.Email
	username := ref.Username
	scene := ref.Scene
	rev := block.Revision
	t1 := lang.NewTime(ref.CreatedAt)
	t2 := lang.NewTime(block.CreatedAt)
	salt := block.Salt

	t1str := strconv.FormatInt(t1.Int(), 10)
	t2str := strconv.FormatInt(t2.Int(), 10)

	inst.makeContent(builder, "email", email.String())
	inst.makeContent(builder, "domain", domain.String())
	inst.makeContent(builder, "username", username)
	inst.makeContent(builder, "scene", scene)
	inst.makeContent(builder, "revision", rev.String())

	inst.makeContent(builder, "length", length)
	inst.makeContent(builder, "charset", charset)
	inst.makeContent(builder, "ref_created_at", t1str)
	inst.makeContent(builder, "block_created_at", t2str)
	inst.makeContent(builder, "salt", salt.String())

	// save to binary-data ...
	str := builder.String()
	data := []byte(str)
	bindata, err := inst.saveBinaryData(data)
	if err != nil {
		return err
	}
	block.Content = bindata.Sum

	vlog.Debug("block.content = [\n %v \n]", str)
	return nil
}

func (inst *innerFastPasswordGenerator) saveBinaryData(data []byte) (*mo.BinaryData, error) {

	ctx := inst.cc
	ser := inst.sc.BinaryDataService
	it1 := new(mo.BinaryData)

	it1.Content = data
	it1.DataType = "block"
	it1.MimeType = "application/x-password-block"

	return ser.Insert(ctx, it1)
}

func (inst *innerFastPasswordGenerator) doStepMakeDigest(step int) error {

	block := &inst.tmpBlock
	content := block.Content // aka. sum of content
	digest := content
	block.Self = dxo.BlockFingerPrint(digest)

	vlog.Debug("block.digest = %v", digest.String())

	return nil
}

func (inst *innerFastPasswordGenerator) doStepMakePath(step int) error {

	block := &inst.tmpBlock
	path := uniformpaths.ComputePathForPasswordBlock(block)
	block.Path = path

	vlog.Debug("block.path = %v", path.String())

	return nil
}

func (inst *innerFastPasswordGenerator) doStepMakePassword(step int) error {

	inst.out.WriteString("todo...")

	return fmt.Errorf("innerFastPasswordGenerator: no impl")
}
