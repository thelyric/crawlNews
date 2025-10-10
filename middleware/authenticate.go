package middleware

import (
	"my-app/common"
	"my-app/component/appctx"
	"my-app/component/tokenprovider"
	"my-app/component/tokenprovider/jwt"
	userstorage "my-app/module/user/storage"

	"github.com/gin-gonic/gin"
)

func getToken(auth string) string {
	if len(auth) > 7 && auth[:7] == "Bearer " {
		return auth[7:]
	}
	return ""
}

func RequireAuth(appContext appctx.AppContext) gin.HandlerFunc {
	jwtProvider := jwt.NewTokenJWTProvider(appContext.SecretKey())

	return func(ctx *gin.Context) {
		token := getToken(ctx.GetHeader("Authorization"))

		if token == "" {
			panic(tokenprovider.ErrInvalidToken)
		}

		payload, err := jwtProvider.Validate(token)
		if err != nil {
			panic(tokenprovider.ErrInvalidToken)
		}

		db := appContext.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		user, err := store.FindDataWithCondition(ctx, map[string]any{
			"id": payload.UserId,
		})

		if err != nil {
			panic(tokenprovider.ErrNotFound)
		}

		if user.Status == 0 {
			panic(common.MakeFailResponse("ERR_USER_DISBLED", "User is disabled"))
		}

		ctx.Set(common.CurrentUser, user)
		ctx.Next()
	}
}
