package biz

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type RealWorld struct {
	Hello string
}

// RealWorldRepo is a Greater repo.
type RealWorldRepo interface {
	Save(context.Context, *RealWorld) (*RealWorld, error)
	Update(context.Context, *RealWorld) (*RealWorld, error)
	FindByID(context.Context, int64) (*RealWorld, error)
	ListByHello(context.Context, string) ([]*RealWorld, error)
	ListAll(context.Context) ([]*RealWorld, error)
}

// GreeterUsecase is a Greeter usecase.
type RealWorldUsecase struct {
	repo RealWorldRepo
	log  *log.Helper
}

// NewRealWorldUsecase new a RealWorld usecase.
func NewRealWorldUsecase(repo RealWorldRepo, logger log.Logger) *RealWorldUsecase {
	return &RealWorldUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// CreateRealWorld creates a RealWorld, and returns the new RealWorld.
func (uc *RealWorldUsecase) CreateRealWorld(ctx context.Context, g *RealWorld) (*RealWorld, error) {
	uc.log.WithContext(ctx).Infof("CreateRealWorld: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
