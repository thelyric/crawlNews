package ginrestaurant

import (
	"my-app/common"
	"my-app/component/appctx"
	restaurantBiz "my-app/module/restaurant/biz"
	restaurantModel "my-app/module/restaurant/model"
	restaurantStorage "my-app/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantModel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewTransportErrorResponse(err))
			return
		}

		user := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = user.GetUserId()

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
