package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
	name := ctx.PostForm("name")
	if !utils.CheckName(name) {
		utils.RespWithParamError(ctx, "name 格式错误")
		return
	}
	err, resp := controller.CtrlReviewPost(mid, ctx.GetInt64("uid"), name, ctx.PostForm("content"), score)
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
	name := ctx.PostForm("name")
	if !utils.CheckName(name) {
		utils.RespWithParamError(ctx, "name 格式错误")
		return
	}
	err, resp := controller.CtrlReviewUpdate(id, ctx.GetInt64("uid"), name, ctx.PostForm("content"), score)
	utils.Resp(ctx, err, resp)
}

func handleReviewStar(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	starList, err := ctx.Cookie("review_star_list")
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
	err, resp := controller.CtrlReviewStar(id, ctx.GetInt64("uid"), value)
	ctx.SetCookie("review_star_list", starList+","+i, 0, "/", "", false, true)
	utils.Resp(ctx, err, resp)
}

func handleReviewBad(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式错误")
		return
	}
	badList, err := ctx.Cookie("review_bad_list")
	i := strconv.FormatInt(id, 10)
	if err != nil {
		badList = ""
	}

	for _, s := range strings.Split(badList, ",") {
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
				Detail:     "已经点踩过了",
			})
			return
		}
	}

	value, err := strconv.ParseBool(ctx.Query("value"))
	if err != nil {
		utils.RespWithParamError(ctx, "value 参数错误")
		return
	}
	err, resp := controller.CtrlReviewBad(id, ctx.GetInt64("uid"), value)
	ctx.SetCookie("review_bad_list", badList+","+i, 0, "/", "", false, true)

	utils.Resp(ctx, err, resp)
}
