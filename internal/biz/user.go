package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	// Implement your password hashing logic here.
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)

	}
	fmt.Printf("hashPassword: %s\n", b)
	return string(b)
}

func verifyPassword(hashed, password string) bool {
	// Implement your password verification logic here.
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	jwtc *conf.JWT
	log *log.Helper
}

// NewRealWorldUsecase new a RealWorld usecase.
func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger, jwtc *conf.JWT) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr,jwtc:jwtc,log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GenerateToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Token, username)
}

// CreateRealWorld creates a RealWorld, and returns the new RealWorld.
func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if _, err := uc.ur.CreateUser(ctx, u); err != nil {
		uc.log.Errorf("failed to create user: %v", err)
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.GenerateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("password is incorrect")
	}

	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.GenerateToken(u.Username),
	}, nil
}
