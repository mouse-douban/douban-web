package middleware

import (
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		err, uid := utils.AuthorizeJWT(accessToken)
		if err != nil {
			utils.AbortWithError(ctx, err)
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}

func WildChecker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		if accessToken != "" { // 不允许外链跳转带 Authorization
			utils.AbortWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40004,
				Info:       "invalid request",
				Detail:     "can not go wild",
			})
			return
		}
	}
}
