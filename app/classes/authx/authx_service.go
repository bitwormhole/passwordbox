package authx

import (
	"context"

	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type Service interface {

	// 登录
	SignIn(cc context.Context, it *dto.Auth) error

	// 注册用户账号
	SignUp(cc context.Context, it *dto.Auth) error
}
