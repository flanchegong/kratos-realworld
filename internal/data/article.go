package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos-layout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Slug        string
	Title       string
	Description string
	Body        string
	CreateAt    time.Time
	UpdateAt    time.Time
	Tags        []Tag `gorm:"many2many:article_tags;"`
}
type Tag struct {
	gorm.Model
	Name string
}

type Fllow struct {
	gorm.Model
	User     uint
	Fllowing uint
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (articleRepo) ListArticle(ctx context.Context, opts ...biz.ListOption) ([]*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (articleRepo) FeedArticle(ctx context.Context, opts ...biz.ListOption) ([]*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}
func (articleRepo) CreateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (articleRepo) GetArticle(ctx context.Context, slug string) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (articleRepo) UpdateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}
func (articleRepo) DeleteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (articleRepo) FavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (articleRepo) UnFavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	// TODO implement me
	panic("implement me")
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
