package middleware

import (
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth0 jwt 必须认证正确
func Auth0() gin.HandlerFunc {
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

// Auth1 可以没有 Authorization，有则认证
func Auth1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		if accessToken == "" {
			ctx.Next()
			return
		}
		Auth0()(ctx)
	}
}
