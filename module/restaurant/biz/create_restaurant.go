package restaurantBiz

import (
	"context"
	"errors"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantModel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantModel.RestaurantCreate) error {
	if data.Name == "" {
		return common.NewBizErrorResponse(errors.New("name canot be empty"))
	}

	if err := biz.store.Create(context, data); err != nil {
		return err
	}

	return nil
}
