package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)
type User struct {
	Username string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewRealWorldUsecase new a RealWorld usecase.
func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}





// CreateRealWorld creates a RealWorld, and returns the new RealWorld.
func (uc *UserUsecase) Register(ctx context.Context,u *User) error {
	if _, err := uc.ur.CreateUser(ctx, u); err != nil {
		uc.log.Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}
