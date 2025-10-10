package ginnews

import (
	"my-app/common"
	newsbiz "my-app/module/news/biz"
	newsmodel "my-app/module/news/model"
	newsrepository "my-app/module/news/repository"
	newsstorage "my-app/module/news/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNews() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data newsmodel.GetArticle

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.NewTransportErrorResponse(err))
		}

		if data.Limit <= 0 {
			data.Limit = 3
		}

		store := newsstorage.NewNewsStore()
		repo := newsrepository.NewNewsRepo(store)
		biz := newsbiz.NewNewsBiz(repo)

		articles, err := biz.GetNews(ctx, &data)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(articles))
	}
}
