package newsrepository

import (
	"context"
	newsmodel "my-app/module/news/model"
)

type newsStore interface {
	FetchLatestNews(ctx context.Context, data *newsmodel.GetArticle) ([]newsmodel.Article, error)
}

type newsRepo struct {
	store newsStore
}

func NewNewsRepo(store newsStore) *newsRepo {
	return &newsRepo{
		store: store,
	}
}

func (s *newsRepo) GetNews(ctx context.Context, data *newsmodel.GetArticle) ([]newsmodel.Article, error) {
	articles, err := s.store.FetchLatestNews(ctx, data)

	if err != nil {
		return nil, err
	}

	return articles, nil
}
