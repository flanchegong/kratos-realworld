package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"
	// Ensure the Article type is defined in the v1 package
	// If not, define it in the appropriate location
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	// // 1. Validate the request
	// if err := validateLoginRequest(req); err != nil {
	// 	return nil, err
	// }

	// // 2. Call the use case to handle the login logic
	// user, err := s.uc.Login(ctx, req.Username, req.Password)
	// if err != nil {
	// 	return nil, err
	// }

	// 3. Create a response and return it
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "flanche",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "flanche",
			// Add other required fields here if necessary
		},
	}, nil
}
