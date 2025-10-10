package restaurantBiz

import (
	"context"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

type GetRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*restaurantModel.Restaurant, error)
	GetRestaurants(context context.Context, filter *restaurantModel.Filter, paging *common.Paging) ([]restaurantModel.Restaurant, error)
}

type getRestaurantsBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantsBiz {
	return &getRestaurantsBiz{
		store: store,
	}
}

func (biz getRestaurantsBiz) GetAllRestaurant(ctx context.Context, filter *restaurantModel.Filter, paging *common.Paging) ([]restaurantModel.Restaurant, error) {
	data, err := biz.store.GetRestaurants(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz getRestaurantsBiz) GetOneRestaurant(context context.Context, id int64) (*restaurantModel.Restaurant, error) {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]any{
		"id":     id,
		"status": 1,
	})

	if err != nil {
		return nil, err
	}

	return oldData, nil
}
