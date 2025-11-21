package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type Passwords struct {
	Base

	Items []*dto.Password `json:"passwords"`
}
