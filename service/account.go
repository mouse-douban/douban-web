package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func RegisterAccountFromUsername(username, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.InsertUser(model.User{
		Username:       username,
		PlaintPassword: password,
	}, dao.UniqueColumnUsername)
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func RegisterAccountFromEmail(email, password string) (err error, accessToken, refreshToken string, uid int64) {
	err, uid = dao.InsertUser(model.User{
		Email:          email,
		PlaintPassword: password,
	}, dao.UniqueColumnEmail)
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
	err, uid = dao.InsertUser(model.User{
		Phone:          phone,
		PlaintPassword: utils.GenerateRandomPassword(),
	}, dao.UniqueColumnPhone)
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}
