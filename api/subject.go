package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

var Tags = utils.StringList{
	"", // 支持空 tag
	"喜剧",
	"生活",
	"爱情",
	"动作",
	"科幻",
	"悬疑",
	"惊悚",
	"动画",
}

func handleSubjectsGet(ctx *gin.Context) {
	var start, limit int

	start, err := strconv.Atoi(ctx.Query("start"))
	if err != nil {
		start = 0
	}

	limit, err = strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 20
	}

	var sort = ctx.Query("sort")
	if sort != "hotest" && sort != "latest" {
		sort = "latest"
	}

	tag, _ := url.QueryUnescape(ctx.Query("tag"))

	if !Tags.Contains(tag) {
		utils.RespWithParamError(ctx, "tag 不支持 "+tag)
		return
	}

	err, resp := controller.CtrlSubjectsGet(start, limit, sort, tag)

	utils.Resp(ctx, err, resp)

}
