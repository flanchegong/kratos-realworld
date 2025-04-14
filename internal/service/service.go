package service

import (
	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
