package ihubproviders

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passwordbox/app/classes/hubproviders"
	"github.com/bitwormhole/passwordbox/app/data/dxo"
	"github.com/bitwormhole/passwordbox/app/web/dto"
)

type ProviderServiceImpl struct {

	//starter:component

	_as func(hubproviders.Service) //starter:as("#")

}

func (inst *ProviderServiceImpl) _impl() hubproviders.Service {
	return inst
}

func (inst *ProviderServiceImpl) Find(ctx context.Context, id dxo.ProviderID) (*dto.Provider, error) {

	return nil, fmt.Errorf("no impl")
}

func (inst *ProviderServiceImpl) ListAll(ctx context.Context) ([]*dto.Provider, error) {

	list := make([]*dto.Provider, 0)

	mock1 := new(dto.Provider)
	mock2 := new(dto.Provider)
	mock3 := new(dto.Provider)

	mock1.Domain = "mock1.example.com"
	mock1.Icon = "http://mock1.example.com/icon.png"
	mock1.WebSite = "http://mock1.example.com/"
	mock1.Label = "mock1"
	mock1.Description = "todo: Description ... "

	mock2.Domain = "mock2.example.com"
	mock2.Icon = "http://mock2.example.com/icon.png"
	mock2.WebSite = "http://mock2.example.com/"
	mock2.Label = "mock2"
	mock2.Description = "todo: Description ... "

	mock3.Domain = "mock3.example.com"
	mock3.Icon = "http://mock3.example.com/icon.png"
	mock3.WebSite = "http://mock3.example.com/"
	mock3.Label = "mock3"
	mock3.Description = "todo: Description ... "

	list = append(list, mock1)
	list = append(list, mock2)
	list = append(list, mock3)

	return list, nil
}

func (inst *ProviderServiceImpl) Insert(ctx context.Context, item *dto.Provider) (*dto.Provider, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ProviderServiceImpl) Update(ctx context.Context, id dxo.ProviderID, item *dto.Provider) (*dto.Provider, error) {
	return nil, fmt.Errorf("no impl")
}

func (inst *ProviderServiceImpl) Remove(ctx context.Context, id dxo.ProviderID) error {
	return fmt.Errorf("no impl")
}
