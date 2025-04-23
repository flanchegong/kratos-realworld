package data

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/biz"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type FollowUser struct {
	gorm.Model
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:100"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	user := User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}

	rv := r.data.db.Create(&user)
	if rv.Error != nil {
		return nil, rv.Error
	}

	createdUser := &biz.User{
		Email:        user.Email,
		Username:     user.Username,
		Bio:          user.Bio,
		Image:        user.Image,
		PasswordHash: user.PasswordHash,
	}
	return createdUser, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	// Implement the logic to get a user by email and return the user and error if any.
	return nil, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	// Implement the logic to get a user by username and return the user and error if any.
	return nil, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// UnFollowUser implements biz.ProfileRepo.
func (r *profileRepo) UnFollowUser(ctx context.Context, username string) (*biz.Profile, error) {
	panic("unimplemented")
}

func (r *profileRepo) GetFollow(ctx context.Context, username string) (*biz.Profile, error) {
	var user User
	err := r.data.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	profile := &biz.Profile{
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
	return profile, nil
}

func (r *profileRepo) GetProfile(ctx context.Context, username string) (*biz.Profile, error) {
	var user User
	err := r.data.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	profile := &biz.Profile{
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
	return profile, nil
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
