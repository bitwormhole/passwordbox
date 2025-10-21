package passwords

import (
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertE2D(src *entity.Password, dst *dto.Password) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.Email = src.Email
	dst.Domain1 = src.Domain1
	dst.Domain2 = src.Domain2
	dst.UserName = src.UserName

	return nil
}

func ConvertD2E(src *dto.Password, dst *entity.Password) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.Email = src.Email
	dst.Domain1 = src.Domain1
	dst.Domain2 = src.Domain2
	dst.UserName = src.UserName

	return nil
}
