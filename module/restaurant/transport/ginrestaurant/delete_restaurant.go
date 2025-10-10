package ginrestaurant

import (
	"errors"
	"my-app/common"
	"my-app/component/appctx"
	restaurantBiz "my-app/module/restaurant/biz"
	restaurantStorage "my-app/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))

		id, err := common.UnmaskID(c.Param("id"))

		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, common.NewTransportErrorResponse(errors.New("invalid ID")))
			return
		}

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}

}
