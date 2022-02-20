package users

import (
	"douban-webend/config"
	"douban-webend/controller"
	"douban-webend/service"
	"douban-webend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		ok = utils.CheckName(account)
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
	case "sms": // + 会转义，使用 %2B 代替
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

	err, resp := controller.CtrlBaseRegister(account, token, kind)
	utils.Resp(ctx, err, resp)
}

func HandleOAuthRedirect(ctx *gin.Context) {
	platform, ok := ctx.GetQuery("platform")
	if !ok {
		utils.AbortWithParamError(ctx, "platform 参数不能为空")
	}
	switch platform {
	case "gitee":
		link := "https://gitee.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code"
		ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf(link, config.Config.GiteeOauthClientId, "https://"+config.Config.ServerIp+"/oauth/gitee"))
	case "github":
		link := "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s"
		ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf(link, config.Config.GithubOauthClientId, "https://"+config.Config.ServerIp+"/oauth/github"))
	default:
		utils.AbortWithParamError(ctx, "不支持这个平台")
		return
	}
}

func HandleOAuthLogin(ctx *gin.Context) {
	platform := ctx.Param("platform")
	if platform != "gitee" && platform != "github" {
		utils.AbortWithParamError(ctx, "不支持这个平台")
		return
	}
	code, ok := ctx.GetQuery("code") // code 不会进入dao层，不需要进行正则检测
	if !ok {
		utils.AbortWithError(ctx, utils.ServerInternalError) // 返回内部错误而不是参数错误
		return
	}
	err, resp := controller.CtrlOAuthLogin(code, platform)
	utils.Resp(ctx, err, resp)
}

func HandleLogin(ctx *gin.Context) {
	kind, ok := ctx.GetPostForm("type")
	if !ok || kind == "" {
		utils.RespWithParamError(ctx, "type 参数不能为空")
		return
	}
	token, ok := ctx.GetPostForm("token")
	if !ok || kind == "" {
		utils.RespWithParamError(ctx, "token 参数不能为空")
		return
	}
	account, ok := ctx.GetPostForm("account")
	if (!ok || account == "") && kind != "refresh" {
		utils.RespWithParamError(ctx, "account 参数不能为空")
		return
	}
	switch kind {
	case "password":
		if !utils.CheckName(account) {
			utils.RespWithParamError(ctx, "用户名格式不支持")
			return
		}
		if !utils.CheckPasswordStrength(token) {
			utils.RespWithParamError(ctx, "密码格式不支持")
			return
		}
	case "email":
		if !utils.MatchEmailFormat(account) {
			utils.RespWithParamError(ctx, "邮箱格式不支持")
			return
		}
		if !utils.CheckPasswordStrength(token) {
			utils.RespWithParamError(ctx, "密码格式不支持")
			return
		}
	case "sms":
		if !utils.MatchPhoneNumber(account) {
			utils.RespWithParamError(ctx, "电话号码格式不支持")
			return
		}
		if !utils.MatchVerifyCode(token) {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40001,
				Info:       "invalid verify code",
				Detail:     "验证码格式有误",
			})
			return
		}
	case "refresh":
		err, uid, tokenType := utils.AuthorizeJWT(token)
		if err != nil {
			utils.RespWithError(ctx, err)
			return
		}
		if tokenType != utils.RefreshTokenType {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40008,
				Info:       "invalid token",
				Detail:     "请不要使用 access_token 来刷新",
			})
			return
		}
		accessToken, refreshToken, err := utils.GenerateTokenPair(uid)
		utils.RespWithData(ctx, utils.RespData{
			HttpStatus: http.StatusOK,
			Status:     20000,
			Info:       "success",
			Data: struct {
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
				Uid          int64  `json:"uid"`
			}{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				Uid:          uid,
			},
		})
		return
	default:
		utils.RespWithParamError(ctx, "type 参数错误, 只能取 password, email, sms, refresh")
		return
	}

	err, resp := controller.CtrlLogin(account, token, kind)
	utils.Resp(ctx, err, resp)
}

func HandleVerify(ctx *gin.Context) {
	kind, ok := ctx.GetQuery("type")
	if !ok || kind == "" {
		utils.RespWithParamError(ctx, "type 参数不能为空")
		return
	}

	target, ok := ctx.GetQuery("value")
	if !ok || kind == "" {
		utils.RespWithParamError(ctx, "value 参数不能为空")
		return
	}

	_, ok = utils.VerifyMap[target]
	if ok {
		utils.RespWithError(ctx, utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40000,
			Info:       "invalid sending",
			Detail:     "请求太频繁",
		})
		return
	}

	switch kind {
	case "sms": // + 号会转译，发请求时使用 %2B
		if !utils.MatchPhoneNumber(target) {
			utils.RespWithParamError(ctx, "value 格式不支持")
			return
		}
		utils.SendRandomVerifyCode("sms", target)
		utils.RespWithDetail(ctx, utils.RespDetail{
			HttpStatus: http.StatusOK,
			Info:       utils.InfoSuccess,
			Status:     20001,
			Data: utils.Detail{
				Detail: "sending sms success",
			},
		})
	case "email":
		if !utils.MatchEmailFormat(target) {
			utils.RespWithParamError(ctx, "value 格式不支持")
			return
		}
		utils.SendRandomVerifyCode("email", target)
		utils.RespWithDetail(ctx, utils.RespDetail{
			HttpStatus: http.StatusOK,
			Info:       utils.InfoSuccess,
			Status:     20002,
			Data: utils.Detail{
				Detail: "sending email success",
			},
		})
	default:
		utils.RespWithParamError(ctx, "type 只能为 email 和 sms")
		return
	}
}

func HandleVerifyAccount(ctx *gin.Context) {
	kind := ctx.Query("type")
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}
	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		utils.RespWithError(ctx, err)
		return
	}
	switch kind {
	case "sms":
		if user.Phone == "" {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40010,
				Info:       "invalid request",
				Detail:     "验证码账户不存在",
			})
			return
		}
		utils.SendRandomVerifyCode("sms", user.Phone)
	case "email":
		if user.Email == "" {
			utils.RespWithError(ctx, utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40010,
				Info:       "invalid request",
				Detail:     "验证码账户不存在",
			})
			return
		}
		utils.SendRandomVerifyCode("email", user.Email)
	default:
		utils.RespWithParamError(ctx, "type 只能为 email 和 sms")
		return
	}
	utils.RespWithDetail(ctx, utils.RespDetail{
		HttpStatus: http.StatusOK,
		Info:       utils.InfoSuccess,
		Status:     20002,
		Data: utils.Detail{
			Detail: "sending " + kind + " success",
		},
	})
}

func HandleForgetPwd(ctx *gin.Context) {
	id := ctx.PostForm("uid")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 格式不支持")
		return
	}
	verify := ctx.PostForm("verify")
	if !utils.MatchVerifyCode(verify) {
		utils.RespWithParamError(ctx, "验证码格式不支持")
		return
	}
	newPwd := ctx.PostForm("new_pwd")
	if !utils.CheckPasswordStrength(newPwd) {
		utils.RespWithParamError(ctx, "新密码格式不支持")
		return
	}
	verifyType := ctx.PostForm("verify_type")
	switch verifyType {
	case "sms", "email":
		err, resp := controller.CtrlResetPwd(uid, verify, verifyType, newPwd)
		utils.Resp(ctx, err, resp)
		return
	default:
		utils.RespWithParamError(ctx, "verify_type 格式不支持")
		return
	}
}
