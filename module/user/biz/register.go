package userbiz

import (
	"context"
	"my-app/common"
	usermodel "my-app/module/user/model"
)

type userCreateStore interface {
	FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*usermodel.User, error)
	CreateUser(context context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type userCreateBiz struct {
	store  userCreateStore
	hasher Hasher
}

func NewRegisterBiz(store userCreateStore, hasher Hasher) *userCreateBiz {
	return &userCreateBiz{
		store:  store,
		hasher: hasher,
	}
}

func (biz userCreateBiz) Resgiter(context context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.FindDataWithCondition(context, map[string]any{
		"email": data.Email,
	})

	if user != nil {
		if user.Status == 0 {
			return common.MakeFailResponse("ERR_USER_DISABLED", "Email đã bị khoá bởi hệ thống!")
		}
		return common.MakeFailResponse("ERR_EMAIL_EXISTED", "Email đã tồn tại!")
	}

	salt, _ := common.GenSalt(32)
	data.Salt = salt
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Role = "User"
	data.Status = 1

	if err := biz.store.CreateUser(context, data); err != nil {
		return err
	}

	return nil
}
