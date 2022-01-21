package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
	"net/http"
)

func RegisterAccountFromUsername(username, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, got := dao.SelectUidFrom("username", username)
	if err == nil || got > 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40005,
			Info:       "invalid request",
			Detail:     "这个账户已经注册了",
		}, "", "", -1
	}
	err, uid = dao.InsertUserFromUserName(model.User{
		Username:       username,
		PlaintPassword: password,
	})
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func RegisterAccountFromEmail(email, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, got := dao.SelectUidFrom("email", email)
	if err == nil || got > 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40005,
			Info:       "invalid request",
			Detail:     "这个账户已经注册了",
		}, "", "", -1
	}
	err, uid = dao.InsertUserFromEmail(model.User{
		Email:          email,
		PlaintPassword: password,
	})
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func RegisterAccountFromSms(phone, verifyCode string) (err error, accessToken, refreshToken string, uid int64) {
	err = utils.VerifyInputCode(phone, "sms", verifyCode)
	if err != nil {
		return
	}
	err, got := dao.SelectUidFrom("phone", phone)
	if err == nil || got > 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40005,
			Info:       "invalid request",
			Detail:     "这个账户已经注册了",
		}, "", "", -1
	}
	err, uid = dao.InsertUserFromPhone(model.User{
		Phone:          phone,
		PlaintPassword: utils.GenerateRandomPassword(),
	})
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func LoginAccountFromGithub(info model.OAuthInfo) (err error, accessToken, refreshToken string, uid int64) {
	panic("TODO")
}

func LoginAccountFromGitee(info model.OAuthInfo) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.SelectUidWithOAuthId(info.OAuthId, info.PlatForm)
	if err != nil || uid <= 0 {
		err, uid = dao.InsertUserFromGiteeId(model.User{
			Username:       info.Username,
			GiteeId:        info.OAuthId,
			PlaintPassword: utils.GenerateRandomPassword(),
			Avatar:         info.Avatar,
		})
		if err != nil {
			return err, "", "", -1
		}
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}
