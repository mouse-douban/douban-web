package dao

import (
	"douban-webend/model"
	"douban-webend/utils"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	UniqueColumnUsername = "username"
	UniqueColumnEmail    = "email"
	UniqueColumnPhone    = "phone"
)

// InsertUser TODO 处理username email phone 的唯一性
func InsertUser(user model.User, uniqueColumn string) (err error, uid int64) {
	var sqlStr string
	switch uniqueColumn {
	case UniqueColumnUsername:
		sqlStr = "INSERT INTO user(username, password) VALUES(?, ?)"
		_, err = dB.Exec(sqlStr, user.Username, user.EncryptPassword())
		if err != nil {
			errStr := fmt.Sprintf("%v", reflect.ValueOf(err))
			if strings.Contains(errStr, "1062") {
				return utils.ServerError{
					HttpStatus: http.StatusBadRequest,
					Status:     40005,
					Info:       "invalid request",
					Detail:     "这个账户已经注册了",
				}, -1
			}
			return
		}
		row := dB.QueryRow("SELECT uid FROM user WHERE username = ?", user.Username)
		err = row.Scan(&uid)
		if err != nil {
			return
		}
	case UniqueColumnEmail:
		sqlStr = "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
		_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Email, user.EncryptPassword())
		if err != nil {
			errStr := fmt.Sprintf("%v", reflect.ValueOf(err))
			if strings.Contains(errStr, "1062") {
				return utils.ServerError{
					HttpStatus: http.StatusBadRequest,
					Status:     40005,
					Info:       "invalid request",
					Detail:     "这个账户已经注册了",
				}, -1
			}
			return
		}
		row := dB.QueryRow("SELECT uid FROM user WHERE email = ?", user.Email)
		err = row.Scan(&uid)
		if err != nil {
			return
		}
	case UniqueColumnPhone:
		sqlStr = "INSERT INTO user(username, phone, password) VALUES(?, ?, ?)"
		_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Phone, user.EncryptPassword())
		if err != nil {
			errStr := fmt.Sprintf("%v", reflect.ValueOf(err))
			if strings.Contains(errStr, "1062") {
				return utils.ServerError{
					HttpStatus: http.StatusBadRequest,
					Status:     40005,
					Info:       "invalid request",
					Detail:     "这个账户已经注册了",
				}, -1
			}
			return
		}
		row := dB.QueryRow("SELECT uid FROM user WHERE phone = ?", user.Phone)
		err = row.Scan(&uid)
		if err != nil {
			return
		}
	}

	return
}
