package service

import (
	"context"

	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateRealWorld(ctx, &biz.RealWorld{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
