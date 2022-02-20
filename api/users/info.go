package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseCommonQueryParams(ctx *gin.Context) (err error, id int64, start, limit int, sort string) {
	idS := ctx.Param("id")
	id, err = strconv.ParseInt(idS, 10, 64)
	paramError := utils.QueryParamError.Copy()
	if err != nil {
		paramError.Info = "id 格式不支持"
		err = paramError
		return
	}
	startS := ctx.Query("start")
	if start, err = strconv.Atoi(startS); startS != "" && err != nil {
		paramError.Info = "start 格式不支持"
		err = paramError
		return
	}
	limitS := ctx.Query("limit")
	if limit, err = strconv.Atoi(limitS); limitS != "" && err != nil {
		paramError.Info = "limit 格式不支持"
		err = paramError
		return
	}
	sort = ctx.Query("sort")
	if sort != "latest" && sort != "hotest" && sort != "" {
		paramError.Info = "sort 格式不支持"
		err = paramError
		return
	}
	if limitS == "" {
		limit = 20
		err = nil
	}
	if startS == "" {
		err = nil
	}
	if sort == "" {
		sort = "hotest"
	}
	return
}

func HandleAccountMovieListGet(ctx *gin.Context) {
	err, uid, start, limit, _ := ParseCommonQueryParams(ctx)
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
	err, uid, start, limit, sort := ParseCommonQueryParams(ctx)
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
	err, uid, start, limit, sort := ParseCommonQueryParams(ctx)
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
