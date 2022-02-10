package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
	"strings"
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
	"奇幻",
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

func handleSubjectGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	mid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	scope := ctx.Query("scope")
	if scope == "" {
		handleSubjectBaseInfoGet(ctx, mid)
		return
	}
	scopes := strings.Split(scope, `,`)
	for _, s := range scopes {
		s = strings.TrimSpace(s)
		if s != "comments" && s != "reviews" && s != "discussions" {
			utils.RespWithParamError(ctx, "scope 格式不支持")
			return
		}
	}
	handleSubjectScopeInfoGet(ctx, mid, scopes)
}

func handleSubjectBaseInfoGet(ctx *gin.Context, mid int64) {
	err, resp := controller.CtrlSubjectBaseInfoGet(mid)
	utils.Resp(ctx, err, resp)
}

func handleSubjectScopeInfoGet(ctx *gin.Context, mid int64, scopes []string) {
	err, resp := controller.CtrlSubjectScopeInfoGet(mid, scopes)
	utils.Resp(ctx, err, resp)
}
