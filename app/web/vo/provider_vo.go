package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type Providers struct {
	Base

	Items []*dto.Provider `json:"providers"`
}
