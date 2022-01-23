package api

import (
	"douban-webend/api/users"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func handleWild(ctx *gin.Context) {
	link, ok := ctx.GetQuery("link")
	match, _ := regexp.Match("http(s?)://.+", []byte(link))
	if !ok {
		utils.AbortWithParamError(ctx, "link 不能为空")
	}
	if !match {
		utils.AbortWithParamError(ctx, "link 参数格式错误, 参考 https://www.baidu.com")
	}
	ctx.Redirect(http.StatusPermanentRedirect, link)
}

func handleMine(ctx *gin.Context) {
	users.HandleAccountBaseInfo(ctx)
}
