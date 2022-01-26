package middleware

import (
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

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

// TLSHandle TLS 证书
func TLSHandle(port string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     port,
		})
		err := secureMiddleware.Process(ctx.Writer, ctx.Request)

		if err != nil {
			return
		}

		ctx.Next()
	}
}
