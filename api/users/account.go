package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"strings"
)

func HandleAccountIndexInfoGet(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}

	ctx.Set("uid", uid)

	scope, ok := ctx.GetQuery("scope")

	if !ok || scope == "" {
		HandleAccountBaseInfoGet(ctx)
		return
	}

	scopes := strings.Split(scope, `,`)

	for _, s := range scopes {
		s = strings.TrimSpace(s)
		if s != "reviews" && s != "before" && s != "after" {
			utils.RespWithParamError(ctx, "scope 格式不支持")
			return
		}
	}
	err, resp := controller.CtrlAccountScopeInfoGet(uid, scope)
	utils.Resp(ctx, err, resp)
}

func HandleAccountBaseInfoGet(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	err, resp := controller.CtrlAccountBaseInfoGet(uid)
	utils.Resp(ctx, err, resp)
}

func HandleAccountInfoUpdate(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")

	scope := ctx.PostForm("scope")

	var params = make(map[string]string)

	for _, s := range strings.Split(scope, ",") {
		s = strings.TrimSpace(s)
		if s != "username" && s != "github_id" && s != "gitee_id" && s != "avatar" && s != "description" {
			utils.AbortWithParamError(ctx, "scope 格式不支持")
			return
		}
		value := ctx.PostForm(s)
		if value == "" {
			utils.AbortWithParamError(ctx, fmt.Sprintf("%v 参数为空", s))
			return
		}
		// 参数检查
		// description 默认会使用预处理，无需检测
		switch s {
		case "username":
			if !utils.CheckName(value) {
				utils.RespWithParamError(ctx, "username 格式不支持")
				return
			}
		case "github_id", "gitee_id":
			if _, err := strconv.Atoi(value); err != nil {
				utils.RespWithParamError(ctx, "oauth id 格式不支持")
				return
			}
		case "avatar":
			if !regexp.MustCompile(`^http(s?)://[0-9a-zA-Z-_./]{0,100}$`).MatchString(value) {
				utils.RespWithParamError(ctx, "avatar 格式不支持")
				return
			}
		}
		params[s] = value
	}

	err, resp := controller.CtrlAccountInfoUpdate(uid, params)
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
		// password 最终会被加密成怪样子，无需 正则检测
		value := ctx.PostForm(s)
		if value == "" {
			utils.AbortWithParamError(ctx, fmt.Sprintf("%v 参数为空", s))
			return
		}
		switch s {
		case "phone":
			if !utils.MatchPhoneNumber(value) {
				utils.AbortWithParamError(ctx, "电话号码格式不支持")
				return
			}
		case "email":
			if !utils.MatchEmailFormat(value) {
				utils.AbortWithParamError(ctx, "邮箱格式不支持")
				return
			}
		}
		params[s] = value
	}

	err, resp := controller.CtrlAccountEXInfoUpdate(uid, params, verifyAccount, verify, verifyType)
	utils.Resp(ctx, err, resp)
}

func HandleAccountDelete(ctx *gin.Context) {
	uid := ctx.GetInt64("uid")
	verify := ctx.PostForm("verify") // 验证码
	if !utils.MatchVerifyCode(verify) {
		utils.AbortWithParamError(ctx, "验证码格式不支持")
		return
	}
	verifyType := ctx.PostForm("verify_type") // 验证方式
	switch verifyType {
	case "sms", "email":
		err, resp := controller.CtrlAccountDelete(uid, verify, verifyType)
		utils.Resp(ctx, err, resp)
		return
	default:
		utils.AbortWithParamError(ctx, "verify_type 格式错误")
		return
	}
}
