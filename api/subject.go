package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type StringList []string

var Tags = StringList{
	"喜剧",
	"生活",
	"爱情",
	"动作",
	"科幻",
	"悬疑",
	"惊悚",
	"动画",
}

func (t StringList) Contains(kind string) bool {
	for _, has := range t {
		if has == kind {
			return true
		}
	}
	return false
}

func handleSubjectsGet(ctx *gin.Context) {
	var start, limit int

	start, err := strconv.Atoi(ctx.PostForm("start"))
	if err != nil {
		start = 0
	}

	limit, err = strconv.Atoi(ctx.PostForm("limit"))
	if err != nil {
		limit = 0
	}

	var sort = ctx.PostForm("sort")
	if sort != "hotest" && sort != "latest" {
		sort = "latest"
	}
	var tag = ctx.PostForm("tag")

	var tags = make([]string, 0)

	for _, s := range strings.Split(tag, ",") {
		s = strings.TrimSpace(s)
		if !Tags.Contains(s) {
			utils.RespWithParamError(ctx, "tag 不支持 "+s)
			return
		}

		tags = append(tags, s)
	}

	err, resp := controller.CtrlSubjectsGet(start, limit, sort, tags)

	utils.Resp(ctx, err, resp)

}
