package userbiz

import (
	"context"
	"my-app/common"
	"my-app/component/tokenprovider"
	usermodel "my-app/module/user/model"
)

type LoginStore interface {
	FindDataWithCondition(context context.Context, condition map[string]any, moreKeys ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store         LoginStore
	hasher        Hasher
	tokenprovider tokenprovider.Provider
	expiry        int
}

func NewLoginBiz(store LoginStore, hasher Hasher, tokenprovider tokenprovider.Provider, expiry int) *loginBiz {
	return &loginBiz{
		store:         store,
		hasher:        hasher,
		tokenprovider: tokenprovider,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]any{
		"email": data.Email,
	})

	if err != nil {
		return nil, common.MakeFailResponse("ERROR_EMAIL_OR_PASSWORD", "Email or password invalid!")
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, common.MakeFailResponse("ERROR_EMAIL_OR_PASSWORD", "Email or password invalid!")
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenprovider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.NewBizErrorResponse(err)
	}

	return accessToken, nil
}
