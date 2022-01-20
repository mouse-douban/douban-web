package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		origin := ctx.GetHeader("Origin")
		if origin == "" {
			origin = "*"
		}
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 处理并放行所有 OPTIONS 请求
		if method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Methods", "PUT, GET, DELETE, POST, PATCH")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, X-Custom-Header")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Max-Age", "1728000")
			ctx.AbortWithStatus(http.StatusNoContent) // 预检成功
			return
		}

		ctx.Next()
	}
}
