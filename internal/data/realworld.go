package data

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type realworldRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRealWorldRepo(data *Data, logger log.Logger) biz.RealWorldRepo {
	return &realworldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realworldRepo) Save(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *realworldRepo) Update(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *realworldRepo) FindByID(context.Context, int64) (*biz.RealWorld, error) {
	return nil, nil
}

func (r *realworldRepo) ListByHello(context.Context, string) ([]*biz.RealWorld, error) {
	return nil, nil
}

func (r *realworldRepo) ListAll(context.Context) ([]*biz.RealWorld, error) {
	return nil, nil
}
