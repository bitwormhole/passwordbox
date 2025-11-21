package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type Examples struct {
	Base

	Items []*dto.Example `json:"examples"`
}
