package controller

import (
	"douban-webend/service"
	"douban-webend/utils"
	"net/http"
)

// CtrlRegister controller 层所有函数均返回 (err error, resp utils.RespData)
func CtrlRegister(account, token, kind string) (err error, resp utils.RespData) {
	err = nil

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
