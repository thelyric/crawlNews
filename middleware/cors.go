package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	whitelist := map[string]bool{
		"https://tieumach.io.vn": true,
		"http://localhost:5173":  true,
	}

	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")

		if whitelist[origin] {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Xử lý preflight request (OPTIONS)
		if ctx.Request.Method == http.MethodOptions {
			if whitelist[origin] {
				ctx.AbortWithStatus(http.StatusNoContent)
				return
			}
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
