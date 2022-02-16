package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
	name := ctx.PostForm("name")
	if !utils.CheckName(name) {
		utils.RespWithParamError(ctx, "name 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionPost(ctx.GetInt64("uid"), mid, name, ctx.PostForm("content"))
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
	name := ctx.PostForm("name")
	if !utils.CheckName(name) {
		utils.RespWithParamError(ctx, "name 格式错误")
		return
	}
	err, resp := controller.CtrlDiscussionUpdate(id, ctx.GetInt64("uid"), name, ctx.PostForm("content"))
	utils.Resp(ctx, err, resp)
}

func handleDiscussionStar(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	starList, err := ctx.Cookie("discussion_star_list")
	i := strconv.FormatInt(id, 10)
	if err != nil {
		starList = ""
	}

	for _, s := range strings.Split(starList, ",") {
		s = strings.TrimSpace(s)
		if s == "" {
			s = "0"
		}
		stared, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			utils.RespWithParamError(ctx, "cookie 错误，请不要擅自修改")
			return
		}
		if stared == id {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: 400,
				Status:     40021,
				Info:       "invalid request",
				Detail:     "已经点赞过了",
			})
			return
		}
	}

	value, err := strconv.ParseBool(ctx.Query("value"))
	if err != nil {
		utils.RespWithParamError(ctx, "value 参数错误")
		return
	}
	err, resp := controller.CtrlDiscussionStar(id, ctx.GetInt64("uid"), value)
	ctx.SetCookie("discussion_star_list", starList+","+i, 0, "/", "", false, true)
	utils.Resp(ctx, err, resp)
}
