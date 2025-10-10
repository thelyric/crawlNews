package restaurantBiz

import (
	"context"
	"errors"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*restaurantModel.Restaurant, error)
	DeleteDataWithCondition(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int64) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]any{
		"id": id,
	})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.NewBizErrorResponse(errors.New("data has been deleted"))
	}

	if id == 0 {
		return common.NewBizErrorResponse(errors.New("id canot be empty"))
	}

	if err := biz.store.DeleteDataWithCondition(context, int(id)); err != nil {
		return err
	}

	return nil
}
