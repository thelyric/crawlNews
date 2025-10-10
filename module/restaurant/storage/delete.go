package restaurantStorage

import (
	"context"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

func (sql *sqlStore) DeleteDataWithCondition(context context.Context, id int) error {
	if err := sql.db.Table(restaurantModel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]any{
		"status": 0,
	}).Error; err != nil {
		return common.NewStorageErrorResponse(err)
	}

	return nil
}
