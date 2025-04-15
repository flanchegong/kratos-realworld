package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Token        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
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

type UserLogin struct {
	Email string `json:"email"`
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
	log *log.Helper
}

// NewRealWorldUsecase new a RealWorld usecase.
func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

// CreateRealWorld creates a RealWorld, and returns the new RealWorld.
func (uc *UserUsecase) Register(ctx context.Context, u *User) error {
	if _, err := uc.ur.CreateUser(ctx, u); err != nil {
		uc.log.Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*User, error) {
	return nil, nil
}
