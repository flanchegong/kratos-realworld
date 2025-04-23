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

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

// NewRealWorldUsecase new a RealWorld usecase.
func NewSocialUsecase(ar ArticleRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
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

// CreateRealWorld creates a RealWorld, and returns the new RealWorld.
func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}
