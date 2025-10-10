package usermodel

import "my-app/common"

const EntityName = "User"

type User struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return int(u.ID)
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID()
}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"-" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *UserCreate) TableName() string {
	return "users"
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (u *UserLogin) TableName() string {
	return "users"
}
