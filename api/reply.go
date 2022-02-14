package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func handleRepliesGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	kind := ctx.Query("type")
	if kind != "review" && kind != "discussion" && kind != "comment" && kind != "reply" {
		utils.RespWithParamError(ctx, "type 格式错误")
		return
	}
	start, err := strconv.Atoi(ctx.Query("start"))
	if err != nil {
		start = 0
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 20
	}

	err, resp := controller.CtrlRepliesGet(id, kind, start, limit)
	utils.Resp(ctx, err, resp)
}

func handleAllRepliesGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}

	err, resp := controller.CtrlAllRepliesGet(id)
	utils.Resp(ctx, err, resp)
}

func handleReplyPost(ctx *gin.Context) {
	idS := ctx.PostForm("pid")
	pid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "pid 格式错误")
		return
	}
	kind := ctx.PostForm("type")
	if kind != "review" && kind != "discussion" && kind != "comment" && kind != "reply" {
		utils.RespWithParamError(ctx, "type 格式错误")
		return
	}

	err, resp := controller.CtrlReplyCreat(ctx.GetInt64("uid"), pid, kind, ctx.PostForm("content"))
	utils.Resp(ctx, err, resp)
}

func handleReplyDelete(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}

	err, resp := controller.CtrlReplyDelete(id, ctx.GetInt64("uid"))
	utils.Resp(ctx, err, resp)
}
