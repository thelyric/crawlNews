package ginuser

import (
	"my-app/common"
	"my-app/component/appctx"
	"my-app/component/hasher"
	"my-app/component/tokenprovider/jwt"
	userbiz "my-app/module/user/biz"
	usermodel "my-app/module/user/model"
	userstorage "my-app/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var loginData usermodel.UserLogin

		if err := ctx.ShouldBind(&loginData); err != nil {
			panic(common.NewTransportErrorResponse(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		tokenprovider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewLoginBiz(store, md5, tokenprovider, 60*60*24*30)

		token, err := biz.Login(ctx, &loginData)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
