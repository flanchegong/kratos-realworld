package data

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	// Implement the logic to create a user and return the created user and error if any.
	return nil
}
 
func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	// Implement the logic to create a user and return the created user and error if any.
	return nil, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
