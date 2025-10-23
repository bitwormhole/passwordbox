package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type Users struct {
	Base

	Items []*dto.User `json:"users"`
}
