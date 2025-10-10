package restaurantStorage

import (
	"context"
	"my-app/common"
	restaurantModel "my-app/module/restaurant/model"
)

func (sql *sqlStore) GetRestaurants(ctx context.Context, filter *restaurantModel.Filter, paging *common.Paging) ([]restaurantModel.Restaurant, error) {
	var restaurants []restaurantModel.Restaurant
	db := sql.db.Table(restaurantModel.Restaurant{}.TableName()).Where("status = ?", 1)

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db.Where("owner_id = ?", f.OwnerId)
		}
	}

	// Phân trang
	if c := paging.FakeCursor; c != "" {
		uuid, err := common.UnmaskID(c)
		if err != nil {
			return nil, common.NewStorageErrorResponse(err)
		}
		db = db.Where("id < ?", uuid)
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db.Offset(offset)
	}

	// Total
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.NewStorageErrorResponse(err)
	}

	result := db.Limit(paging.Limit).Order("id desc").Find(&restaurants)
	if result.Error != nil {
		return nil, common.NewStorageErrorResponse(result.Error)
	}

	if len(restaurants) > 0 {
		last := restaurants[len(restaurants)-1]
		last.Mask(false)
		if last.FakeId != nil {
			paging.NextCursor = *last.FakeId
		} else {
			paging.NextCursor = ""
		}
	}

	return restaurants, nil
}
