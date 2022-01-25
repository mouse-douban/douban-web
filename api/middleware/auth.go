package middleware

import (
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		err, uid, kind := utils.AuthorizeJWT(accessToken)
		if err != nil {
			utils.AbortWithError(ctx, err)
			return
		}
		if kind != utils.AccessTokenType {
			utils.AbortWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40008,
				Info:       "invalid token",
				Detail:     "请不要使用 refresh_token 认证",
			})
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}
