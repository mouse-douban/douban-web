package controller

import (
	"douban-webend/config"
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
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
		account = strings.Replace(account, "%2B", "+", -1) // 替换 url_encode
		err, accessToken, refreshToken, uid = service.RegisterAccountFromSms(account, token)
	}

	if err != nil {
		return err, utils.RespData{}
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
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
		Info:       utils.InfoSuccess,
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
		// 获取 token
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
		// 获取信息
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
		// 注册｜登录
		err, accessToken, refreshToken, uid = service.LoginAccountFromGitee(info)
	case "github":
		// 获取 token
		postUrl := fmt.Sprintf(GithubToken, config.Config.GithubOauthClientId, config.Config.GithubOauthClientSecret, code)
		tokenChan := utils.GetPOSTBytesWithEmptyBody(postUrl)

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

		// 注册｜登录
		info.PlatForm = platform
		err, accessToken, refreshToken, uid = service.LoginAccountFromGithub(info)
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
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

func CtrlAccountBaseInfoGet(uid int64) (err error, resp utils.RespData) {

	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       user,
	}
	return
}

// CtrlAccountInfoUpdate 需要对 params 内的参数做正则检测
func CtrlAccountInfoUpdate(uid int64, params map[string]string) (err error, resp utils.RespData) {
	err = service.UpdateUserInfo(uid, params)
	if err != nil {
		return
	}
	return nil, utils.NoDetailSuccessResp
}

// CtrlAccountEXInfoUpdate 需要对 params 内的参数做正则检测
func CtrlAccountEXInfoUpdate(uid int64, params map[string]string, verifyAccount, verifyCode, verifyType string) (err error, resp utils.RespData) {
	err = utils.VerifyInputCode(verifyAccount, verifyType, verifyCode)
	if err != nil {
		return
	}
	var kind = verifyType
	if verifyType == "sms" {
		kind = "phone"
	}
	params[kind] = verifyAccount // sms | phone 替换
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

func CtrlAccountScopeInfoGet(uid int64, scopes string) (err error, resp utils.RespData) {

	var data = make(map[string]interface{})
	for _, s := range strings.Split(scopes, `,`) {
		s = strings.TrimSpace(s)
		err = service.GetAccountSnapshots(uid, s, &data)
		if err != nil {
			return
		}
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       data,
	}
	return
}

func CtrlAccountMovieListGet(uid int64, start, limit int) (err error, resp utils.RespData) {
	movieList := make([]model.MovieList, 0)
	err = service.GetAccountMovieList(uid, &movieList, start, limit)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       movieList,
	}
	return
}

func CtrlAccountBeforeGet(uid int64, start, limit int, sort string) (err error, resp utils.RespData) {
	before := make([]model.Comment, 0)
	err = service.GetAccountComments(uid, "before", &before, start, limit, sort)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       before,
	}
	return
}

func CtrlAccountAfterGet(uid int64, start, limit int, sort string) (err error, resp utils.RespData) {
	after := make([]model.Comment, 0)
	err = service.GetAccountComments(uid, "after", &after, start, limit, sort)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       after,
	}
	return
}

func CtrlAccountReviewSnapshotsGet(uid int64, start, limit int, sort string) (err error, resp utils.RespData) {
	reviews := make([]model.ReviewSnapshot, 0)
	err = service.GetAccountReviewSnapshots(uid, &reviews, start, limit, sort)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       reviews,
	}
	return
}

var followerSync = sync.Mutex{}

func CtrlAccountFollow(uid, id int64, kind string) (err error, resp utils.RespData) {
	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return
	}
	if kind == "users" && id == uid {
		return utils.ServerError{
			HttpStatus: http.StatusOK,
			Status:     22222,
			Info:       utils.InfoSuccess,
			Detail:     "你每时每刻都在关注着自己",
		}, utils.RespData{}
	}
	switch kind {
	case "users":
		// 检查用户是否已经关注了
		if check, ok := user.Following.Users[id]; ok && check {
			return utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40011,
				Info:       "invalid request",
				Detail:     "已经关注了这个id",
			}, utils.RespData{}
		}
		// 检查用户是否存在
		err, _ := service.GetAccountBaseInfo(id)
		if err != nil {
			return err, utils.RespData{}
		}
		// 关注用户
		user.Following.Users[id] = true
		bytes, err := json.Marshal(user.Following.Users)
		if err != nil {
			return err, utils.RespData{}
		}
		err = service.UpdateUserInfo(uid, map[string]string{"following_users": string(bytes)})
		if err != nil {
			return err, utils.RespData{}
		}
	case "lists":
		if check, ok := user.Following.Lists[id]; ok && check {
			return utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40011,
				Info:       "invalid request",
				Detail:     "已经关注了这个id",
			}, utils.RespData{}
		}

		followerSync.Lock() // 锁起来
		// 修改 movie list 的 followers
		err, list := service.GetMovieList(id)
		if err != nil { // 找不到
			followerSync.Unlock() // 解锁
			return err, utils.RespData{}
		}

		list.Followers += 1
		err = service.UpdateMovieListInfo(id, map[string]interface{}{"followers": list.Followers}, false)
		followerSync.Unlock() // 解锁

		if err != nil {
			return err, utils.RespData{}
		}

		user.Following.Lists[id] = true
		bytes, err := json.Marshal(user.Following.Lists)
		if err != nil {
			return err, utils.RespData{}
		}
		err = service.UpdateUserInfo(uid, map[string]string{"following_lists": string(bytes)})
		if err != nil {
			return err, utils.RespData{}
		}
	}
	return nil, utils.NoDetailSuccessResp
}

func CtrlAccountUnfollow(uid, id int64, kind string) (err error, resp utils.RespData) {
	err, user := service.GetAccountBaseInfo(uid)
	if err != nil {
		return
	}
	if kind == "users" && id == uid {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     44444,
			Info:       "invalid request",
			Detail:     "你每时每刻都在关注着自己",
		}, utils.RespData{}
	}
	switch kind {
	case "users":
		if check, ok := user.Following.Users[id]; !ok || !check {
			return utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40012,
				Info:       "invalid request",
				Detail:     "无法取关没有关注的id",
			}, utils.RespData{}
		}
		delete(user.Following.Users, id)
		bytes, err := json.Marshal(user.Following.Users)
		if err != nil {
			return err, utils.RespData{}
		}
		err = service.UpdateUserInfo(uid, map[string]string{"following_users": string(bytes)})
		if err != nil {
			return err, utils.RespData{}
		}
	case "lists":
		if check, ok := user.Following.Lists[id]; !ok || !check {
			return utils.ServerError{
				HttpStatus: http.StatusBadRequest,
				Status:     40012,
				Info:       "invalid request",
				Detail:     "无法取关没有关注的id",
			}, utils.RespData{}
		}

		followerSync.Lock() // 锁起来
		// 修改 movie list 的 followers
		err, list := service.GetMovieList(id)
		if err != nil { // 找不到
			followerSync.Unlock() // 解锁
			return err, utils.RespData{}
		}

		list.Followers -= 1
		err = service.UpdateMovieListInfo(id, map[string]interface{}{"followers": list.Followers}, false)
		followerSync.Unlock() // 解锁

		if err != nil {
			return err, utils.RespData{}
		}

		delete(user.Following.Lists, id)
		bytes, err := json.Marshal(user.Following.Lists)
		if err != nil {
			return err, utils.RespData{}
		}
		err = service.UpdateUserInfo(uid, map[string]string{"following_lists": string(bytes)})
		if err != nil {
			return err, utils.RespData{}
		}
	}
	return nil, utils.NoDetailSuccessResp
}

const ImgBucketUrl = "https://image-1259160349.cos.ap-chengdu.myqcloud.com"

func CtrlAvatarUpload(uid int64, img io.Reader, name string) (err error, resp utils.RespData) {
	var url = ImgBucketUrl
	uuid := utils.GenerateRandomUUID()
	path := "/avatar/" + uuid + "." + strings.Split(name, `.`)[1]
	url += path
	err = service.UpdateUserInfo(uid, map[string]string{"avatar": url})
	utils.UploadFile(ImgBucketUrl, path, img)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       map[string]string{"avatar": url},
	}
	return
}
