package ginuser

import (
	"my-app/common"
	"my-app/component/appctx"
	userbiz "my-app/module/user/biz"
	userstorage "my-app/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfo(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ginCtx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		user := ginCtx.MustGet(common.CurrentUser).(common.Requester)

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewUserInfoBiz(store)

		data, err := biz.UserInfo(ginCtx, user.GetUserId())

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		ginCtx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}

}
