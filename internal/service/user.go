package service

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/errors"

	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"
	// Ensure the Article type is defined in the v1 package
	// If not, define it in the appropriate location
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHTTPError(422, "email can't be empty", "validation error")
	}

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "flanche",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,

			// Add other required fields here if necessary
		},
	}, nil
}
