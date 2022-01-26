package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func HandleAccountIndexInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}

	ctx.Set("uid", uid)

	scope, ok := ctx.GetQuery("scope")

	if !ok || scope == "" {
		HandleAccountBaseInfo(ctx)
		return
	}

	scopes := strings.Split(scope, `,`)

	for _, s := range scopes {
		s = strings.TrimSpace(s)
		if s != "reviews" && s != "movie_list" && s != "before" && s != "after" {
			utils.RespWithParamError(ctx, "scope 格式不支持")
			return
		}
	}

	ctx.Set("scope", scope)
	handleAccountScopeInfo(ctx)
}

func HandleAccountBaseInfo(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	err, resp := controller.CtrlAccountBaseInfo(uid)
	utils.Resp(ctx, err, resp)
}

func handleAccountScopeInfo(ctx *gin.Context) {

}

func HandleUserBefore(ctx *gin.Context) {

}

func HandleUserAfter(ctx *gin.Context) {

}

func HandleAccountInfoUpdate(ctx *gin.Context) {
	_, ok := ctx.Get("uid")
	verify := ctx.PostForm("verify")          // 验证码
	verifyType := ctx.PostForm("verify_type") // 验证方式
	verifyAccount := ctx.PostForm("verify_account")

	if !ok && (verify == "" || verifyType == "" || verifyAccount == "") {
		utils.AbortWithParamError(ctx, "无法通过 jwt 认证且验证码认证信息为空")
		return
	}

	switch verifyType {
	case "sms":
		if !utils.MatchPhoneNumber(verifyAccount) {
			utils.AbortWithParamError(ctx, "电话号码格式不支持")
			return
		}
		if !utils.MatchVerifyCode(verify) {
			utils.AbortWithParamError(ctx, "验证码格式不支持")
			return
		}
	case "email":
		if !utils.MatchEmailFormat(verifyAccount) {
			utils.AbortWithParamError(ctx, "邮箱格式不支持")
			return
		}
		if !utils.MatchVerifyCode(verify) {
			utils.AbortWithParamError(ctx, "验证码格式不支持")
			return
		}
	default:
		if ok {
			break
		}
		utils.AbortWithParamError(ctx, "verify_type 格式错误")
		return
	}

	scope := ctx.PostForm("scope")

	var params = make(map[string]string)

	for _, s := range strings.Split(scope, ",") {
		s = strings.TrimSpace(s)
		if s != "username" && s != "github_id" && s != "gitee_id" && s != "avatar" {
			utils.AbortWithParamError(ctx, "scope 格式不支持")
			return
		}
		value := ctx.PostForm(s)
		if value == "" {
			utils.AbortWithParamError(ctx, fmt.Sprintf("%v 参数为空", s))
			return
		}
		params[s] = value
	}
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}

	err, resp := controller.CtrlAccountInfoUpdate(uid, params, !ok, verifyAccount, verify, verifyType)
	utils.Resp(ctx, err, resp)

}

func HandleAccountEXInfoUpdate(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")

	verify := ctx.PostForm("verify")          // 验证码
	verifyType := ctx.PostForm("verify_type") // 验证方式
	verifyAccount := ctx.PostForm("verify_account")

	switch verifyType {
	case "sms":
		if !utils.MatchPhoneNumber(verifyAccount) {
			utils.AbortWithParamError(ctx, "电话号码格式不支持")
			return
		}
		if !utils.MatchVerifyCode(verify) {
			utils.AbortWithParamError(ctx, "验证码格式不支持")
			return
		}
	case "email":
		if !utils.MatchEmailFormat(verifyAccount) {
			utils.AbortWithParamError(ctx, "邮箱格式不支持")
			return
		}
		if !utils.MatchVerifyCode(verify) {
			utils.AbortWithParamError(ctx, "验证码格式不支持")
			return
		}
	default:
		utils.AbortWithParamError(ctx, "verify_type 格式错误")
		return
	}

	scope := ctx.PostForm("scope")

	var params = make(map[string]string)

	for _, s := range strings.Split(scope, ",") {
		s = strings.TrimSpace(s)
		if s != "password" && s != "email" && s != "phone" {
			utils.AbortWithParamError(ctx, "scope 格式不支持")
			return
		}
		value := ctx.PostForm(s)
		if value == "" {
			utils.AbortWithParamError(ctx, fmt.Sprintf("%v 参数为空", s))
			return
		}
		params[s] = value
	}

	err, resp := controller.CtrlAccountEXInfoUpdate(uid, params, verifyAccount, verify, verifyType)
	utils.Resp(ctx, err, resp)
}
