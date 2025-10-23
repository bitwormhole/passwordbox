package vo

import "github.com/bitwormhole/passwordbox/app/web/dto"

type KeyPairs struct {
	Base

	Items []*dto.KeyPair `json:"keypairs"`
}
