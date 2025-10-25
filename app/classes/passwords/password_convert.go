package passwords

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/uniformpaths"
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertE2D(src1 *entity.PasswordChain, src2 *entity.PasswordBlock, dst *dto.Password) error {

	if dst == nil {
		return fmt.Errorf("param: dst is nil")
	}
	if src1 == nil {
		src1 = new(entity.PasswordChain)
	}
	if src2 == nil {
		src2 = new(entity.PasswordBlock)
	}

	dst.ID = src1.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src1.BaseEntity, &dst.BaseDTO)

	dst.Email = src1.Email
	dst.Domain = src1.Domain
	dst.UserName = src1.Username
	dst.Scene = src1.Scene

	dst.Path = src1.Path

	dst.Revision = src2.Revision
	dst.Charset = src2.Charset
	dst.Length = src2.Length
	dst.Salt = src2.Salt

	return nil
}

func ConvertD2E(src *dto.Password, dst1 *entity.PasswordChain, dst2 *entity.PasswordBlock) error {

	if src == nil {
		src = new(dto.Password)
	}
	if dst1 == nil {
		dst1 = new(entity.PasswordChain)
	}
	if dst2 == nil {
		dst2 = new(entity.PasswordBlock)
	}

	up := uniformpaths.ComputePathForPassword(src)

	dst1.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst1.BaseEntity)

	dst1.Email = src.Email
	dst1.Domain = src.Domain
	dst1.Username = src.UserName
	dst1.Scene = src.Scene

	dst2.Revision = src.Revision
	dst2.Charset = src.Charset
	dst2.Length = src.Length
	dst2.Salt = src.Salt

	dst1.Path = up
	dst2.Path = up

	return nil
}
