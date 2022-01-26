package controller

import (
	"douban-webend/config"
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CtrlBaseRegister controller 层所有函数均返回 (err error, resp utils.RespData)
func CtrlBaseRegister(account, token, kind string) (err error, resp utils.RespData) {

	var accessToken, refreshToken string
	var uid int64

	switch kind {
	case "password":
		err, accessToken, refreshToken, uid = service.RegisterAccountFromUsername(account, token)
	case "email":
		err, accessToken, refreshToken, uid = service.RegisterAccountFromEmail(account, token)
	case "sms":
		err, accessToken, refreshToken, uid = service.RegisterAccountFromSms(account, token)
	}

	if err != nil {
		return err, utils.RespData{}
	}

	resp = utils.RespData{
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
	}
	return
}

func CtrlLogin(account, token, kind string) (err error, resp utils.RespData) {

	var accessToken, refreshToken string
	var uid int64

	switch kind {
	case "password":
		err, accessToken, refreshToken, uid = service.LoginAccountFromUsername(account, token)
	case "email":
		err, accessToken, refreshToken, uid = service.LoginAccountFromEmail(account, token)
	case "sms":
		err, accessToken, refreshToken, uid = service.LoginAccountFromSms(account, token)
	}

	if err != nil {
		return err, utils.RespData{}
	}

	resp = utils.RespData{
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
	}

	return
}

const (
	GithubToken       = "https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s"
	GithubOpenAPIUser = "https://api.github.com/user"

	GiteeToken       = "https://gitee.com/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s"
	GiteeRedirectUri = "https://%s/oauth/gitee"
	GiteeOpenAPIUser = "https://gitee.com/api/v5/user/"
)

func CtrlOAuthLogin(code, platform string) (err error, resp utils.RespData) {

	err = nil

	var accessToken, refreshToken string
	var uid int64

	var info model.OAuthInfo

	switch platform {
	case "gitee":
		postUrl := fmt.Sprintf(GiteeToken, code, config.Config.GiteeOauthClientId, fmt.Sprintf(GiteeRedirectUri, config.Config.ServerIp), config.Config.GiteeOauthClientSecret)
		tokenChan := utils.GetPOSTBytesWithEmptyBody(postUrl) // 请求token
		var token struct {
			AccessToken string `json:"access_token"`
		}
		tokenJson := <-tokenChan
		if len(tokenJson) == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}
		err = json.Unmarshal(tokenJson, &token)
		if err != nil {
			return
		}
		infoCh := utils.GetGETBytes(GiteeOpenAPIUser+"?access_token="+token.AccessToken, nil)
		infoJson := <-infoCh
		if len(infoJson) == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}
		err = json.Unmarshal(infoJson, &info)
		if err != nil || info.OAuthId == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}
		info.PlatForm = platform
		err, accessToken, refreshToken, uid = service.LoginAccountFromGitee(info)
	case "github":
		postUrl := fmt.Sprintf(GithubToken, config.Config.GithubOauthClientId, config.Config.GithubOauthClientSecret, code)
		tokenChan := utils.GetPOSTBytesWithEmptyBody(postUrl) // 请求 token

		tokenB := <-tokenChan
		if len(tokenB) == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}
		token := strings.Split(strings.Split(string(tokenB), "=")[1], "&")[0]

		infoCh := utils.GetGETBytes(GithubOpenAPIUser, map[string]string{ // 请求数据
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": "token " + token,
		})

		infoJson := <-infoCh

		if len(infoJson) == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}
		err = json.Unmarshal(infoJson, &info)
		if err != nil || info.OAuthId == 0 {
			return utils.ServerInternalError, utils.RespData{}
		}

		info.PlatForm = platform
		err, accessToken, refreshToken, uid = service.LoginAccountFromGithub(info)
	}

	resp = utils.RespData{
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
	}

	return
}

func CtrlAccountBaseInfo(uid int64) (err error, resp utils.RespData) {

	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40009,
			Info:       "invalid request",
			Detail:     "没有这个账户",
		}, utils.RespData{}
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       user,
	}
	return
}

func CtrlAccountInfoUpdate(uid int64, params map[string]string) (err error, resp utils.RespData) {
	err = service.UpdateUserInfo(uid, params)
	if err != nil {
		return
	}
	return nil, utils.NoDetailSuccessResp
}

func CtrlAccountEXInfoUpdate(uid int64, params map[string]string, verifyAccount, verifyCode, verifyType string) (err error, resp utils.RespData) {
	err = utils.VerifyInputCode(verifyAccount, verifyType, verifyCode)
	if err != nil {
		return
	}
	var kind = verifyType
	if verifyType == "sms" {
		kind = "phone"
	}
	params[kind] = verifyAccount // 替换
	err = service.UpdateUserInfo(uid, params)
	if err != nil {
		return
	}
	return nil, utils.NoDetailSuccessResp
}

func CtrlResetPwd(uid int64, verifyCode, verifyType, newPwd string) (err error, resp utils.RespData) {
	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return
	}
	var account string
	switch verifyType {
	case "sms":
		account = user.Phone
	case "email":
		account = user.Email
	}
	if account == "" {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40010,
			Info:       "invalid request",
			Detail:     "验证码账户不存在",
		}, utils.RespData{}
	}

	err = utils.VerifyInputCode(account, verifyType, verifyCode)
	if err != nil {
		return
	}
	err = service.UpdateUserInfo(uid, map[string]string{"password": newPwd})
	if err != nil {
		return
	}
	return nil, utils.NoDetailSuccessResp
}

func CtrlAccountDelete(uid int64, verifyCode, verifyType string) (err error, resp utils.RespData) {
	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return
	}
	var account string
	switch verifyType {
	case "sms":
		account = user.Phone
	case "email":
		account = user.Email
	}
	if account == "" {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40010,
			Info:       "invalid request",
			Detail:     "验证码账户不存在",
		}, utils.RespData{}
	}

	err = utils.VerifyInputCode(account, verifyType, verifyCode)
	if err != nil {
		return
	}

	err = service.DeleteUser(uid)
	if err != nil {
		return
	}

	return nil, utils.NoDetailSuccessResp
}
