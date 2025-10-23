package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type Authx struct {
	Base

	Items []*dto.Auth `json:"auth"`
}
