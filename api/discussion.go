package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"encoding/json"
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
	value, err := strconv.ParseBool(ctx.Query("value"))
	if err != nil {
		utils.RespWithParamError(ctx, "value 参数错误")
		return
	}

	starListStr, err := ctx.Cookie("discussion_star_list")

	if err != nil {
		starListStr = ""
	}
	starListStr = "[" + starListStr + "]"

	var starList []int64

	err = json.Unmarshal([]byte(starListStr), &starList)
	if err != nil {
		utils.RespWithParamError(ctx, "cookie 错误，请不要擅自修改")
		return
	}

	var rmIndex = -1
	for index, stared := range starList {
		if stared == id && value {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: 400,
				Status:     40021,
				Info:       "invalid request",
				Detail:     "已经点赞过了",
			})
			return
		}
		if stared == id && !value {
			rmIndex = index
		}

	}

	if value {
		starList = append(starList, id)
	}

	if rmIndex == -1 && !value {
		utils.RespWithError(ctx, utils.ServerError{
			HttpStatus: 400,
			Status:     40021,
			Info:       "invalid request",
			Detail:     "已经取消点赞过了",
		})
		return
	}

	if rmIndex != -1 {
		starList = append(starList[:rmIndex], starList[rmIndex+1:]...)
	}

	starListB, err := json.Marshal(starList)
	if err != nil {
		utils.RespWithError(ctx, utils.ServerInternalError)
		return
	}
	starListStr = strings.TrimRight(strings.TrimLeft(string(starListB), "["), "]")

	err, resp := controller.CtrlDiscussionStar(id, ctx.GetInt64("uid"), value)
	ctx.SetCookie("discussion_star_list", starListStr, 0, "/", "", false, true)
	utils.Resp(ctx, err, resp)
}
