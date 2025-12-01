package frontend

import (
	"github.com/bitwormhole/passwordbox/core/pemapi"
)

func GetAPI() pemapi.API {
	fe := GetFrontend()
	return fe
}
