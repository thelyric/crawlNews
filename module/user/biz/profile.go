package userbiz

import (
	"context"
	"errors"
	"my-app/common"
	usermodel "my-app/module/user/model"
)

type userInfoStore interface {
	FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*usermodel.User, error)
}

type userInfoBiz struct {
	store userInfoStore
}

func NewUserInfoBiz(store userCreateStore) *userInfoBiz {
	return &userInfoBiz{
		store: store,
	}
}

func (biz userInfoBiz) UserInfo(context context.Context, id int) (*usermodel.User, error) {
	user, _ := biz.store.FindDataWithCondition(context, map[string]any{
		"id": id,
	})
	if user == nil {
		return nil, common.NewRecordNotFoundResponse(errors.New("User không tồn tại!"))
	}

	return user, nil
}
