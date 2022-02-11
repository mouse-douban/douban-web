package api

import (
	"douban-webend/api/users"
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func handleWild(ctx *gin.Context) {
	link, ok := ctx.GetQuery("link")
	match, _ := regexp.Match(`^http(s?)://[0-9a-zA-Z-_./]{0,100}$`, []byte(link))
	if !ok {
		utils.AbortWithParamError(ctx, "link 不能为空")
	}
	if !match {
		utils.AbortWithParamError(ctx, "link 参数格式错误, 参考 https://www.baidu.com")
	}
	ctx.Redirect(http.StatusPermanentRedirect, link)
}

func handleMine(ctx *gin.Context) {
	users.HandleAccountBaseInfoGet(ctx)
}

func handleAvatarUpload(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	img, err := ctx.FormFile("img")
	if err != nil {
		utils.RespWithError(ctx, err)
		return
	}
	file, err := img.Open()
	if err != nil {
		utils.RespWithError(ctx, err)
		return
	}
	err, resp := controller.CtrlAvatarUpload(uid, file, img.Filename)
	utils.Resp(ctx, err, resp)
}
