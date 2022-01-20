package users

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRegister(ctx *gin.Context) {
	kind, ok := ctx.GetPostForm("type")
	if !ok || kind == "" {
		utils.RespWithParamError(ctx, "type 参数不能为空")
		return
	}
	account, ok := ctx.GetPostForm("account")
	if !ok || account == "" {
		utils.RespWithParamError(ctx, "account 参数不能为空")
		return
	}
	token, ok := ctx.GetPostForm("token")
	if !ok || token == "" {
		utils.RespWithParamError(ctx, "token 参数不能为空")
		return
	}

	switch kind {
	case "password":
		ok = utils.CheckPasswordStrength(token)
		if !ok {
			utils.RespWithParamError(ctx, "密码格式不支持")
			return
		}
		ok = utils.CheckUsername(account)
		if !ok {
			utils.RespWithParamError(ctx, "用户名称格式不支持")
			return
		}
	case "email":
		ok = utils.CheckPasswordStrength(token)
		if !ok {
			utils.RespWithParamError(ctx, "密码格式不支持")
			return
		}
		ok = utils.MatchEmailFormat(account)
		if !ok {
			utils.RespWithParamError(ctx, "邮箱格式不对")
			return
		}
	case "sms":
		ok = utils.MatchPhoneNumber(account)
		if !ok {
			utils.RespWithParamError(ctx, "电话格式不支持")
			return
		}

		ok = utils.MatchVerifyCode(token)
		if !ok {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40001,
				Info:       "invalid verify code",
				Detail:     "验证码格式有误",
			})
			return
		}
	default:
		utils.RespWithParamError(ctx, "type 参数错误, 只能取 password, email, sms")
		return
	}

	err, resp := controller.CtrlRegister(account, token, kind)
	if err != nil {
		utils.RespWithError(ctx, err)
		return
	}
	utils.RespWithData(ctx, resp)

}

func HandleLogin(ctx *gin.Context) {

}
