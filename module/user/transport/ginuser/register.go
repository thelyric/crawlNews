package ginuser

import (
	"my-app/common"
	"my-app/component/appctx"
	"my-app/component/hasher"
	userbiz "my-app/module/user/biz"
	usermodel "my-app/module/user/model"
	userstorage "my-app/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resgiter(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ginCtx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := ginCtx.ShouldBind(&data); err != nil {
			panic(common.NewTransportErrorResponse(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Resgiter(ginCtx, &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ginCtx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}

}
