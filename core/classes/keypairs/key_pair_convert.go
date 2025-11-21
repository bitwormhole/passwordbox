package keypairs

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

func ConvertE2D(src *entity.KeyPair, dst *dto.KeyPair) error {
	return fmt.Errorf("no impl")
}

func ConvertD2E(src *dto.KeyPair, dst *entity.KeyPair) error {
	return fmt.Errorf("no impl")
}
