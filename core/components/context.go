package components

import "github.com/bitwormhole/passwordbox/core/data/properties"

////////////////////////////////////////////////////////////////////////////////

type Context interface {
	GetEnvironment() properties.Table

	GetProperties() properties.Table

	GetComponents() ComponentManager
}

////////////////////////////////////////////////////////////////////////////////

type ContextFactory interface {
	Create(b *ContextBuilder) (Context, error)
}

////////////////////////////////////////////////////////////////////////////////

type innerContextFactory struct{}

// Create implements ContextFactory.
func (inst *innerContextFactory) Create(cb *ContextBuilder) (Context, error) {
	ctx := new(innerComponentContext)
	ctx2 := ctx.init()

	inst.innerCopyProps(cb, ctx2)
	inst.innerCopyEnv(cb, ctx2)
	inst.innerCopyCom(cb, ctx2)

	err := inst.makeIndexer(ctx)
	if err != nil {
		return nil, err
	}

	return ctx2, nil
}

func (inst *innerContextFactory) innerCopyProps(src *ContextBuilder, dst Context) {

	tmp := src.innerGetData().properties.Export(nil)
	dst.GetProperties().Import(tmp)

}

func (inst *innerContextFactory) innerCopyEnv(src *ContextBuilder, dst Context) {

	tmp := src.innerGetData().environment.Export(nil)
	dst.GetEnvironment().Import(tmp)

}

func (inst *innerContextFactory) innerCopyCom(src *ContextBuilder, dst Context) {

	from := src.innerGetData().registry.ListAll()
	to := dst.GetComponents().GetRegistry()

	for _, cr := range from {
		to.Register(cr)
	}

}

func (inst *innerContextFactory) makeIndexer(ctx *innerComponentContext) error {

	ids := ctx.registry.ListIDs()
	cache := ctx.cache
	dst := ctx.indexer

	for _, id := range ids {
		h, err := cache.fetchByID(id)
		if err != nil {
			return err
		}
		dst.put(h)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
