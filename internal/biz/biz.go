package biz

import (
	"time"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSocialUsecase, NewUserUsecase)

type Article struct {
	Slug        string
	Title       string
	Description string
	Body        string
	CreateAt    time.Time
	UpdateAt    time.Time
	Tags        []string
}

