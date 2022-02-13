package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func handleDiscussionGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionGet(id)
	utils.Resp(ctx, err, resp)
}

func handleDiscussionPost(ctx *gin.Context) {
	midS := ctx.PostForm("mid")
	mid, err := strconv.ParseInt(midS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "mid 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionPost(ctx.GetInt64("uid"), mid, ctx.PostForm("name"), ctx.PostForm("content"))
	utils.Resp(ctx, err, resp)
}

func handleDiscussionDelete(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionDelete(id, ctx.GetInt64("uid"))
	utils.Resp(ctx, err, resp)
}

func handleDiscussionUpdate(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionUpdate(id, ctx.GetInt64("uid"), ctx.PostForm("name"), ctx.PostForm("content"))
	utils.Resp(ctx, err, resp)
}
