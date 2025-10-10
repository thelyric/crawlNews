package restaurantModel

import "my-app/common"

type Restaurant struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Addr    string         `json:"addr" gorm:"column:addr;"`
	OwnerID string         `json:"owner_id" gorm:"column:owner_id;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID()
}

type RestaurantCreate struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Addr    string         `json:"addr" gorm:"column:addr;"`
	OwnerID string         `json:"owner_id" gorm:"column:owner_id;"`
	UserId  int            `json:"-" gorm:"column:user_id;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID()
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}
