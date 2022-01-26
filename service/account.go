package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
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

func LoginAccountFromUsername(username, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.SelectUidFrom("username", username)
	if err != nil || uid <= 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40006,
			Info:       "invalid request",
			Detail:     "这个账户没有注册",
		}, "", "", -1
	}
	err, encrypt := dao.SelectEncryptPassword(uid)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password))
	if err != nil {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40007,
			Info:       "invalid request",
			Detail:     "密码错误",
		}, "", "", -1
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

func LoginAccountFromEmail(email, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.SelectUidFrom("email", email)
	if err != nil || uid <= 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40006,
			Info:       "invalid request",
			Detail:     "这个账户没有注册",
		}, "", "", -1
	}
	err, encrypt := dao.SelectEncryptPassword(uid)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(password))
	if err != nil {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40007,
			Info:       "invalid request",
			Detail:     "密码错误",
		}, "", "", -1
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

func LoginAccountFromSms(phone, verifyCode string) (err error, accessToken, refreshToken string, uid int64) {
	err = utils.VerifyInputCode(phone, "sms", verifyCode)
	if err != nil {
		return
	}
	err, uid = dao.SelectUidFrom("phone", phone)
	if err != nil || uid <= 0 {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40006,
			Info:       "invalid request",
			Detail:     "这个账户没有注册",
		}, "", "", -1
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func LoginAccountFromGithub(info model.OAuthInfo) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.SelectUidWithOAuthId(info.OAuthId, info.PlatForm)
	if err != nil || uid <= 0 {
		err, uid = dao.InsertUserFromGithubId(model.User{
			Username:       info.Username,
			GithubId:       info.OAuthId,
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

func GetAccountBaseInfo(uid int64) (err error, user model.User) {
	err, user = dao.SelectBaseUserInfo(uid)
	user.Phone = strings.Replace(user.Phone, "%2B", "+", -1) // + 转义
	if !utils.MatchPhoneNumber(user.Phone) {                 // 排除 UUID 占位
		user.Phone = ""
	}
	if !utils.MatchEmailFormat(user.Email) { // 排除 UUID 占位
		user.Email = ""
	}
	return
}

func UpdateUserInfo(uid int64, params map[string]string) (err error) {
	for key, value := range params {
		if key == "password" { // 加密
			user := model.User{PlaintPassword: value}
			value = user.EncryptPassword()
		}
		err = dao.RawUpdateUserInfo(uid, key, value)
		if err != nil {
			return
		}
	}
	return
}
