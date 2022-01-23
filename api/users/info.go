package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func HandleAccountIndexInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}

	ctx.Set("uid", uid)

	scope, ok := ctx.GetQuery("scope")

	if !ok || scope == "" {
		HandleAccountBaseInfo(ctx)
		return
	}

	scopes := strings.Split(scope, `,`)

	for _, s := range scopes {
		if s != "reviews" && s != "movie_list" && s != "before" && s != "after" {
			utils.RespWithParamError(ctx, "scope 格式不支持")
			return
		}
	}

	ctx.Set("scope", scope)
	handleAccountScopeInfo(ctx)
}

func HandleAccountBaseInfo(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	err, resp := controller.CtrlAccountBaseInfo(uid)
	utils.Resp(ctx, err, resp)
}

func handleAccountScopeInfo(ctx *gin.Context) {

}

func HandleUserBefore(ctx *gin.Context) {

}

func HandleUserAfter(ctx *gin.Context) {

}
