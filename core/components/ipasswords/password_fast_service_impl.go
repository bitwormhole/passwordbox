package ipasswords

import (
	"context"
	"sync"
	"time"

	"github.com/bitwormhole/passwordbox/app/classes/passwords"
	"github.com/bitwormhole/passwordbox/app/web/dto"
	"github.com/starter-go/vlog"
)

type PasswordFastServiceImpl struct {

	//starter:component

	_as func(passwords.FastService) //starter:as("#")

	DaoRefs   passwords.ChainDAO //starter:inject("#")
	DaoBlocks passwords.BlockDAO //starter:inject("#")

	mutex sync.Mutex
}

func (inst *PasswordFastServiceImpl) _impl() passwords.FastService {
	return inst
}

func (inst *PasswordFastServiceImpl) MakeFastGen(ctx context.Context, item *dto.Password) (*dto.Password, error) {

	mtx := &inst.mutex
	mtx.Lock()
	defer mtx.Unlock()
	time.Sleep(time.Second * 2) // 保护

	gen := new(innerFastPasswordGenerator)
	gen.InitWithDTO(item)
	plain, err := gen.Generate()
	if err != nil {
		return nil, err
	}

	txt := plain.String()
	vlog.Info("fast-gen: %v", txt)

	return item, nil
}
