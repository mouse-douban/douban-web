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

// TODO Redis 缓存
// 点赞一般是高频访问的接口，直接对数据库IO 效率非常低下
// 在用户量巨大的情况下就简单的点赞功能会引出以下问题
// · 用户的点赞关系如何储存？
// · 数据如何归档入库？
// · 如何支持高峰期大量的并发请求？
// · 用户点赞记录存满了怎么办？
// · 分布式存储怎么合理调度
// · ...
func handleStar(ctx *gin.Context) {
	ctx.Query("type")
}

// TODO Redis 缓存
func handleBad(ctx *gin.Context) {

}

func handleSwagger(ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, "https://douban.skygard.cn/swagger/openapi.json")
}

func handleSearch(ctx *gin.Context) {
	words := ctx.Query("key")
	if words == "" {
		utils.RespWithParamError(ctx, "key 不能为空")
		return
	}
	err, resp1 := controller.CtrlWhatCelebritiesNameLike(words)
	if err != nil {
		utils.RespWithError(ctx, utils.ServerInternalError)
		return
	}
	err, resp2 := controller.CtrlWhatSubjectsNameLike(words)
	if err != nil {
		utils.RespWithError(ctx, utils.ServerInternalError)
		return
	}
	response := make([]utils.RespData, 2)
	response[0] = resp1
	response[1] = resp2
	ctx.JSON(http.StatusOK, response)
}
