package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func handleCommentGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlCommentGet(id)
	utils.Resp(ctx, err, resp)
}

func handleCommentPost(ctx *gin.Context) {
	midS := ctx.PostForm("mid")
	mid, err := strconv.ParseInt(midS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "mid 格式错误")
		return
	}
	scoreS := ctx.PostForm("score")
	score, err := strconv.Atoi(scoreS)
	if err != nil || score > 5 || score < 1 {
		utils.RespWithParamError(ctx, "score 格式错误")
		return
	}
	kind := ctx.PostForm("type")
	if kind != "after" && kind != "before" {
		utils.RespWithParamError(ctx, "type 格式错误")
		return
	}
	err, resp := controller.CtrlCommentPost(mid, ctx.GetInt64("uid"), ctx.PostForm("content"), time.Now(), score, strings.Split(ctx.PostForm("tag"), ","), kind, 0)
	utils.Resp(ctx, err, resp)
}

func handleCommentDelete(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlCommentDelete(id, ctx.GetInt64("uid"))
	utils.Resp(ctx, err, resp)
}

func handleCommentUpdate(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	scoreS := ctx.PostForm("score")
	score, err := strconv.Atoi(scoreS)
	if err != nil || score > 5 || score < 1 {
		utils.RespWithParamError(ctx, "score 格式错误")
		return
	}
	err, resp := controller.CtrlCommentUpdate(id, ctx.GetInt64("uid"), strings.Split(ctx.PostForm("tag"), ","), ctx.PostForm("content"), score)
	utils.Resp(ctx, err, resp)
}
