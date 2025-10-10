package restaurantStorage

import (
	"context"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

func (sql *sqlStore) Create(context context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.NewStorageErrorResponse(err)
	}

	return nil
}
