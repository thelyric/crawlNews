package userstorage

import (
	"context"
	"my-app/common"
	usermodel "my-app/module/user/model"

	"gorm.io/gorm"
)

func (sql *sqlStore) FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User

	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NewRecordNotFoundResponse(err)
		}

		return nil, common.NewStorageErrorResponse(err)
	}

	return &data, nil
}
