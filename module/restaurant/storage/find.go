package restaurantStorage

import (
	"context"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"

	"gorm.io/gorm"
)

func (sql *sqlStore) FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*restaurantModel.Restaurant, error) {
	var data restaurantModel.Restaurant

	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewRecordNotFoundResponse(err)
		}
		return nil, common.NewStorageErrorResponse(err)
	}

	return &data, nil
}
