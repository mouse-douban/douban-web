package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func parseCommonQueryParams(ctx *gin.Context) (err error, uid int64, start, limit int, sort string) {
	id := ctx.Param("id")
	uid, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}
	startS := ctx.Query("start")
	if start, err = strconv.Atoi(startS); startS != "" && err != nil {
		utils.RespWithParamError(ctx, "start 格式不支持")
		return
	}
	limitS := ctx.Query("limit")
	if limit, err = strconv.Atoi(limitS); limitS != "" && err != nil {
		utils.RespWithParamError(ctx, "limit 格式不支持")
		return
	}
	sort = ctx.Query("sort")
	if sort != "latest" && sort != "hotest" && sort != "" {
		utils.RespWithParamError(ctx, "sort 格式不支持")
		err = utils.ServerInternalError
		return
	}
	return
}

func HandleAccountMovieListGet(ctx *gin.Context) {
	err, uid, start, limit, _ := parseCommonQueryParams(ctx)
	if err != nil {
		return
	}
	if limit == 0 {
		limit = 10
	}
	err, resp := controller.CtrlAccountMovieListGet(uid, start, limit)
	utils.Resp(ctx, err, resp)
}

func HandleAccountBeforeGet(ctx *gin.Context) {
	handleAccountCommentGet(ctx, "before")
}

func HandleAccountAfterGet(ctx *gin.Context) {
	handleAccountCommentGet(ctx, "after")
}

func handleAccountCommentGet(ctx *gin.Context, kind string) {
	err, uid, start, limit, sort := parseCommonQueryParams(ctx)
	if err != nil {
		return
	}
	if limit == 0 {
		limit = 10
	}
	if sort == "" {
		sort = "latest"
	}
	switch kind {
	case "before":
		err, resp := controller.CtrlAccountBeforeGet(uid, start, limit, sort)
		utils.Resp(ctx, err, resp)
	case "after":
		err, resp := controller.CtrlAccountAfterGet(uid, start, limit, sort)
		utils.Resp(ctx, err, resp)
	}
}

func HandleAccountReviewsGet(ctx *gin.Context) {
	err, uid, start, limit, sort := parseCommonQueryParams(ctx)
	if err != nil {
		return
	}
	if limit == 0 {
		limit = 10
	}
	if sort == "" {
		sort = "latest"
	}
	err, resp := controller.CtrlAccountReviewSnapshotsGet(uid, start, limit, sort)
	utils.Resp(ctx, err, resp)
}

func HandleAccountFollow(ctx *gin.Context) {
	handleFollows(ctx, "follow")
}

func HandleAccountUnFollow(ctx *gin.Context) {
	handleFollows(ctx, "unfollow")
}

func handleFollows(ctx *gin.Context, which string) {
	uidS := ctx.Param("id")
	uid, err := strconv.ParseInt(uidS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "uid 格式不支持")
		return
	}
	idS := ctx.Query("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}
	kind := ctx.Query("type")
	switch kind {
	case "users", "lists":
		switch which {
		case "follow":
			err, resp := controller.CtrlAccountFollow(uid, id, kind)
			utils.Resp(ctx, err, resp)
		case "unfollow":
			err, resp := controller.CtrlAccountUnfollow(uid, id, kind)
			utils.Resp(ctx, err, resp)
		}
	default:
		utils.RespWithParamError(ctx, "type 格式不支持")
	}

}
