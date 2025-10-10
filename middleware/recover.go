package middleware

import (
	"fmt"
	"my-app/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Recover(appContext appctx.AppContext) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				ctx.Header("Content-Type", "application/json")
// 				fmt.Println("error", err)

// 				if appErr, ok := err.(*common.FailRes); ok {
// 					ctx.AbortWithStatusJSON(http.StatusBadRequest, appErr)
// 					panic(err)
// 				}

// 				appErr := common.NewInternalErrorResponse(err.(error))
// 				ctx.AbortWithStatusJSON(http.StatusInternalServerError, appErr)
// 				panic(err)
// 			}
// 		}()
// 		ctx.Next()
// 	}
// }

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")
				fmt.Println("error", err)

				if appErr, ok := err.(*common.FailRes); ok {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, appErr)
					panic(err)
				}

				appErr := common.NewInternalErrorResponse(err.(error))
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, appErr)
				panic(err)
			}
		}()
		ctx.Next()
	}
}
