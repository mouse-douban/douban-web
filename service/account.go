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
	if err != nil {
		err = utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40009,
			Info:       "invalid request",
			Detail:     "没有这个账户",
		}
		return
	}
	user.Phone = strings.Replace(user.Phone, "%2B", "+", -1) // + 转义
	if !utils.MatchPhoneNumber(user.Phone) {                 // 排除 UUID 占位
		user.Phone = ""
	}
	if !utils.MatchEmailFormat(user.Email) { // 排除 UUID 占位
		user.Email = ""
	}
	return
}

func GetAccountSnapshots(uid int64, scope string, user *map[string]interface{}) (err error) {
	switch scope {
	case "reviews":
		reviews := make([]model.ReviewSnapshot, 0)
		err = GetAccountReviewSnapshots(uid, &reviews, 0, 6, "latest")
		if err != nil {
			return
		}
		(*user)["reviews"] = reviews
	case "before", "after":
		comments := make([]model.Comment, 0)
		err = GetAccountComments(uid, scope, &comments, 0, 10, "latest")
		if err != nil {
			return
		}
		(*user)[scope] = comments
	default:
		return utils.ServerInternalError
	}
	return nil
}

var orderBys = map[string]string{
	"latest": "date DESC",
	"hotest": "stars DESC",
}

func GetAccountReviewSnapshots(uid int64, data *[]model.ReviewSnapshot, start, limit int, sort string) (err error) {
	err, *data = dao.SelectUserReviewSnapshot(uid, orderBys[sort], start, limit)
	return
}

func GetAccountComments(uid int64, kind string, data *[]model.Comment, start, limit int, sort string) (err error) {
	err, *data = dao.SelectUserComments(uid, kind, orderBys[sort], start, limit)
	return
}

func GetAccountMovieList(uid int64, data *[]model.MovieList, start, limit int) (err error) {
	err, *data = dao.SelectUserMovieList(uid, start, limit)
	return
}

// UpdateUserInfo 执行多条更新请求，加入事务处理，需要对 params 做好正则检测
func UpdateUserInfo(uid int64, params map[string]string) (err error) {
	tx, err := dao.OpenTransaction() // 开启一个事务
	if err != nil {
		return
	}
	for key, value := range params {
		if key == "password" { // 加密
			user := model.User{PlaintPassword: value}
			value = user.EncryptPassword()
		}
		if key == "description" { // 预防 SQL 注入
			value = utils.ReplaceXSSKeywords(value)
			value = utils.ReplaceWildUrl(value) // 换掉外链，前端就可以渲染这个外链了
			err = dao.UpdateUserDescription(uid, value, tx)
			if err != nil {
				dao.RollBackTransaction(tx)
				return
			}
			continue
		}
		err = dao.RawUpdateUserInfo(uid, key, value, tx)
		if err != nil {
			dao.RollBackTransaction(tx)
			return
		}
	}
	dao.CommitTransaction(tx)
	return
}

func DeleteUser(uid int64) (err error) {
	// todo 先删除 user 的所有子表 使用事务
	return dao.DeleteUser(uid)
}
