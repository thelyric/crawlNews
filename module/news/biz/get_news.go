package newsbiz

import (
	"context"
	newsmodel "my-app/module/news/model"
)

type newsRepo interface {
	GetNews(ctx context.Context, data *newsmodel.GetArticle) ([]newsmodel.Article, error)
}

type newsBiz struct {
	repo newsRepo
}

func NewNewsBiz(repo newsRepo) *newsBiz {
	return &newsBiz{
		repo: repo,
	}
}

func (biz *newsBiz) GetNews(ctx context.Context, data *newsmodel.GetArticle) ([]newsmodel.Article, error) {
	articles, err := biz.repo.GetNews(ctx, data)

	if err != nil {
		return nil, err
	}

	return articles, nil
}
