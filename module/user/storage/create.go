package userstorage

import (
	"context"
	"my-app/common"
	usermodel "my-app/module/user/model"
)

func (sql *sqlStore) CreateUser(context context.Context, data *usermodel.UserCreate) error {
	db := sql.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.NewStorageErrorResponse(err)

	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.NewStorageErrorResponse(err)
	}
	return nil
}
