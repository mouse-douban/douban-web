package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func handleMovieListGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	lid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlMovieListGet(lid)
	utils.Resp(ctx, err, resp)
}

func handleMovieListCreate(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	listS := ctx.PostForm("list")
	list := make([]int64, 0)
	for _, s := range strings.Split(listS, ",") {
		s = strings.TrimSpace(s)
		mid, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			utils.RespWithParamError(ctx, "list 格式错误")
			return
		}
		list = append(list, mid)
	}
	err, resp := controller.CtrlMovieListCreate(uid, ctx.PostForm("name"), ctx.PostForm("description"), list)
	utils.Resp(ctx, err, resp)
}

func handleMovieListDelete(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	idS := ctx.Param("id")
	lid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	err, resp := controller.CtrlDeleteMovieList(lid, uid)
	utils.Resp(ctx, err, resp)
}

func handleMovieListUpdate(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	idS := ctx.Param("id")
	lid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	scopeS := ctx.PostForm("scope")
	params := make(map[string]interface{})
	for _, s := range strings.Split(scopeS, ",") {
		s = strings.TrimSpace(s)
		if s != "name" && s != "description" {
			utils.RespWithParamError(ctx, "scope 格式错误")
			return
		}
		value := ctx.PostForm(s)
		if value == "" {
			utils.RespWithParamError(ctx, s+" 不能为空")
			return
		}
		params[s] = value
	}
	err, resp := controller.CtrlUpdateMovieList(lid, uid, params)
	utils.Resp(ctx, err, resp)
}

func handleMovieListMovieAdd(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	idS := ctx.Param("id")
	lid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	midS := ctx.PostForm("mid")
	newMids := make([]int64, 0)
	for _, s := range strings.Split(midS, ",") {
		s = strings.TrimSpace(s)
		mid, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			utils.RespWithParamError(ctx, "mid 格式错误")
			return
		}
		newMids = append(newMids, mid)
	}
	err, resp := controller.CtrlMovieListMovieAdd(lid, uid, newMids)
	utils.Resp(ctx, err, resp)
}

func handleMovieListMovieRemove(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	idS := ctx.Param("id")
	lid, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	midS := ctx.PostForm("mid")
	removeMids := make([]int64, 0)
	for _, s := range strings.Split(midS, ",") {
		s = strings.TrimSpace(s)
		mid, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			utils.RespWithParamError(ctx, "mid 格式错误")
			return
		}
		removeMids = append(removeMids, mid)
	}
	err, resp := controller.CtrlMovieListMovieRemove(lid, uid, removeMids)
	utils.Resp(ctx, err, resp)
}
