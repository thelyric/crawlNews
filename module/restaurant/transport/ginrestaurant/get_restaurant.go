package ginrestaurant

import (
	"errors"
	"my-app/common"
	"my-app/component/appctx"
	restaurantBiz "my-app/module/restaurant/biz"
	restaurantModel "my-app/module/restaurant/model"
	restaurantStorage "my-app/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.NewTransportErrorResponse(err))
		}

		paging.Fullfill()

		var filter restaurantModel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.NewTransportErrorResponse(err))
		}

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBiz.GetRestaurantStore(store)

		data, err := biz.GetRestaurants(c, &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}

func GetOneRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := common.UnmaskID(context.Param("id"))

		if err != nil || id == 0 {
			panic(common.NewTransportErrorResponse(errors.New("invalid ID")))
		}

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBiz.NewGetRestaurantBiz(store)

		data, err := biz.GetOneRestaurant(context, id)

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
