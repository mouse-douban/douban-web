package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func handleReviewGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlReviewGet(id)
	utils.Resp(ctx, err, resp)
}

func handleReviewPost(ctx *gin.Context) {
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
	err, resp := controller.CtrlReviewPost(mid, ctx.GetInt64("uid"), ctx.PostForm("name"), ctx.PostForm("content"), score)
	utils.Resp(ctx, err, resp)
}

func handleReviewDelete(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlReviewDelete(id, ctx.GetInt64("uid"))
	utils.Resp(ctx, err, resp)
}

func handleReviewUpdate(ctx *gin.Context) {
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
	err, resp := controller.CtrlReviewUpdate(id, ctx.GetInt64("uid"), ctx.PostForm("name"), ctx.PostForm("content"), score)
	utils.Resp(ctx, err, resp)
}
