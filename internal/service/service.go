package service

import（
	"github.com/google/wire"
	v1 "github.com/go-kratos/kratos-layout/api/realworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

） 

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
	log *log.Helper
}


// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.RealWorldUsecase，logger log.logger) *RealWorldService {
	return &RealWorldService{uc: uc,log:log.NewHelper(logger)}
}




