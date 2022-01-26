package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"fmt"
)

// SelectUidFrom 只允许内部调用，不会出现sql注入
func SelectUidFrom(accountType, account string) (err error, uid int64) {
	sqlStr := "SELECT uid FROM user WHERE " + accountType + " = '" + account + "'"
	err = dB.QueryRow(sqlStr).Scan(&uid)
	return
}

func SelectEncryptPassword(uid int64) (err error, encrypt string) {
	sqlStr := "SELECT password FROM user WHERE uid = ?"
	err = dB.QueryRow(sqlStr, uid).Scan(&encrypt)
	return
}

func SelectUserSnapshot(uid int64, scope []string) (err error, user model.User) {
	panic("TODO")
}

func SelectUidWithOAuthId(oauthID int64, platform string) (err error, uid int64) {
	sqlStr := "SELECT uid FROM user WHERE " + platform + "_id = ?"
	err = dB.QueryRow(sqlStr, oauthID).Scan(&uid)
	return
}

func SelectBaseUserInfo(uid int64) (err error, user model.User) {
	user.GithubId = -1
	user.GiteeId = -1
	sqlStr := "SELECT username, uid, github_id, gitee_id, email, phone, avatar FROM user WHERE uid = ?"
	var githubId sql.NullInt64 // 可能为 null
	var giteeId sql.NullInt64  // 可能为 null
	err = dB.QueryRow(sqlStr, uid).Scan(
		&user.Username,
		&user.Uid,
		&githubId,
		&giteeId,
		&user.Email,
		&user.Phone,
		&user.Avatar,
	)
	if githubId.Valid {
		user.GithubId = githubId.Int64
	}
	if giteeId.Valid {
		user.GiteeId = giteeId.Int64
	}
	return
}

func InsertUserFromUserName(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, password, email, phone) VALUES(?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 phone email
	_, err = dB.Exec(sqlStr, user.Username, user.EncryptPassword(), utils.GenerateRandomUUID(), utils.GenerateRandomUUID())
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE username = ?", user.Username)
	err = row.Scan(&uid)
	return
}

func InsertUserFromEmail(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, email, password, phone) VALUES(?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 phone
	_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Email, user.EncryptPassword(), utils.GenerateRandomUUID())
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE email = ?", user.Email)
	err = row.Scan(&uid)
	return
}

func InsertUserFromPhone(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, phone, password, email) VALUES(?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 email
	_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Phone, user.EncryptPassword(), utils.GenerateRandomUUID())
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE phone = ?", user.Phone)
	err = row.Scan(&uid)
	return
}

func InsertUserFromGiteeId(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, phone, password, email, gitee_id, avatar) VALUES(?, ?, ?, ?, ?, ?)"
	_, err = dB.Exec(sqlStr, user.Username+utils.GenerateRandomUUID()[:8], utils.GenerateRandomUUID(), user.EncryptPassword(), utils.GenerateRandomUUID(), user.GiteeId, user.Avatar)
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE gitee_id = ?", user.GiteeId)
	err = row.Scan(&uid)
	return
}

func InsertUserFromGithubId(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, phone, password, email, github_id, avatar) VALUES(?, ?, ?, ?, ?, ?)"
	_, err = dB.Exec(sqlStr, user.Username+utils.GenerateRandomUUID()[:8], utils.GenerateRandomUUID(), user.EncryptPassword(), utils.GenerateRandomUUID(), user.GithubId, user.Avatar)
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE github_id = ?", user.GithubId)
	err = row.Scan(&uid)
	return
}

// RawUpdateUserInfo 不建议引用
func RawUpdateUserInfo(uid int64, which, what string) (err error) {
	sqlStr := fmt.Sprintf("UPDATE user SET %s=? WHERE uid = ?", which)
	_, err = dB.Exec(sqlStr, what, uid)
	return
}
