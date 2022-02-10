package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"encoding/json"
	"fmt"
	"strings"
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

func SelectUserReviewSnapshot(uid int64, orderBy string) (err error, reviews []model.ReviewSnapshot) {
	sqlStr := "SELECT r.id, r.uid, r.name, r.score, r.date, r.stars, r.bads, r.reply_cnt, r.content, u.avatar, u.username, r.mid FROM review r JOIN user u ON r.uid = ? AND u.uid = ? ORDER BY " + orderBy
	rows, err := dB.Query(sqlStr, uid, uid)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerWarning("rows 关闭异常", err)
		}
	}(rows)

	for rows.Next() {
		var review model.ReviewSnapshot
		err = rows.Scan(
			&review.Id,
			&review.Uid,
			&review.Name,
			&review.Score,
			&review.Date,
			&review.Stars,
			&review.Bads,
			&review.ReplyCnt,
			&review.Brief,
			&review.Avatar,
			&review.Username,
			&review.Mid,
		)
		// 修短成 165 个中文
		var end = 165 * 3
		if end > len(review.Brief) {
			end = len(review.Brief)
		}
		review.Brief = review.Brief[:end]
		if err != nil {
			return
		}
		reviews = append(reviews, review)
	}
	return
}

func SelectUserMovieList(uid int64) (err error, list []model.MovieList) {
	sqlStr := "SELECT id, uid, name, date, followers, list, description FROM movie_list WHERE uid = ? ORDER BY followers DESC"
	rows, err := dB.Query(sqlStr, uid)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerWarning("rows 关闭异常", err)
		}
	}(rows)

	for rows.Next() {
		var movieList model.MovieList
		var listJson string
		err = rows.Scan(
			&movieList.Id,
			&movieList.Uid,
			&movieList.Name,
			&movieList.Date,
			&movieList.Followers,
			&listJson,
			&movieList.Description,
		)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(listJson), &movieList.List)
		if err != nil {
			return
		}
		list = append(list, movieList)
	}
	return
}

func SelectUserComments(uid int64, kind string, orderBy string) (err error, comments []model.Comment) {
	sqlStr := "SELECT c.id, c.uid, c.content, c.date, c.score, c.tag, c.type, c.stars, u.username, c.mid FROM comment c JOIN user u ON c.uid = ? AND u.uid = ? AND c.type = ? ORDER BY " + orderBy
	rows, err := dB.Query(sqlStr, uid, uid, kind)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerWarning("rows 关闭异常", err)
		}
	}(rows)

	for rows.Next() {
		var comment model.Comment
		var tag string
		err = rows.Scan(
			&comment.Id,
			&comment.Uid,
			&comment.Content,
			&comment.Date,
			&comment.Score,
			&tag,
			&comment.Type,
			&comment.Stars,
			&comment.Username,
			&comment.Mid,
		)
		if err != nil {
			return
		}
		comment.Tag = strings.Split(tag, ",")
		comments = append(comments, comment)
	}
	return
}

func SelectUidWithOAuthId(oauthID int64, platform string) (err error, uid int64) {
	sqlStr := "SELECT uid FROM user WHERE " + platform + "_id = ?"
	err = dB.QueryRow(sqlStr, oauthID).Scan(&uid)
	return
}

func SelectBaseUserInfo(uid int64) (err error, user model.User) {
	user.GithubId = -1
	user.GiteeId = -1
	sqlStr := "SELECT username, uid, github_id, gitee_id, email, phone, avatar, description, following_users, following_lists FROM user WHERE uid = ?"
	var githubId sql.NullInt64 // 可能为 null
	var giteeId sql.NullInt64  // 可能为 null
	var followingUserJson string
	var followingListJson string
	err = dB.QueryRow(sqlStr, uid).Scan(
		&user.Username,
		&user.Uid,
		&githubId,
		&giteeId,
		&user.Email,
		&user.Phone,
		&user.Avatar,
		&user.Description,
		&followingUserJson,
		&followingListJson,
	)
	if githubId.Valid {
		user.GithubId = githubId.Int64
	}
	if giteeId.Valid {
		user.GiteeId = giteeId.Int64
	}
	err = json.Unmarshal([]byte(followingUserJson), &user.Following.Users)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(followingListJson), &user.Following.Lists)
	if err != nil {
		return
	}
	return
}

func InsertUserFromUserName(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, password, email, phone, following_users, following_lists) VALUES(?, ?, ?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 phone email
	_, err = dB.Exec(sqlStr, user.Username, user.EncryptPassword(), utils.GenerateRandomUUID(), utils.GenerateRandomUUID(), "{}", "{}")
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE username = ?", user.Username)
	err = row.Scan(&uid)
	return
}

func InsertUserFromEmail(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, email, password, phone, following_users, following_lists) VALUES(?, ?, ?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 phone
	_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Email, user.EncryptPassword(), utils.GenerateRandomUUID(), "{}", "{}")
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE email = ?", user.Email)
	err = row.Scan(&uid)
	return
}

func InsertUserFromPhone(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, email, password, phone, following_users, following_lists) VALUES(?, ?, ?, ?, ?, ?)"
	// 使用 UUID 来初始化非空唯一键 email
	_, err = dB.Exec(sqlStr, utils.GenerateRandomUserName(), user.Phone, user.EncryptPassword(), utils.GenerateRandomUUID(), "{}", "{}")
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE phone = ?", user.Phone)
	err = row.Scan(&uid)
	return
}

func InsertUserFromGiteeId(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, email, password, phone, following_users, following_lists) VALUES(?, ?, ?, ?, ?, ?)"
	_, err = dB.Exec(sqlStr, user.Username+utils.GenerateRandomUUID()[:8], utils.GenerateRandomUUID(), user.EncryptPassword(), utils.GenerateRandomUUID(), user.GiteeId, user.Avatar, "{}", "{}")
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE gitee_id = ?", user.GiteeId)
	err = row.Scan(&uid)
	return
}

func InsertUserFromGithubId(user model.User) (err error, uid int64) {
	sqlStr := "INSERT INTO user(username, email, password, phone, following_users, following_lists) VALUES(?, ?, ?, ?, ?, ?)"
	_, err = dB.Exec(sqlStr, user.Username+utils.GenerateRandomUUID()[:8], utils.GenerateRandomUUID(), user.EncryptPassword(), utils.GenerateRandomUUID(), user.GithubId, user.Avatar, "{}", "{}")
	if err != nil {
		return
	}
	row := dB.QueryRow("SELECT uid FROM user WHERE github_id = ?", user.GithubId)
	err = row.Scan(&uid)
	return
}

// RawUpdateUserInfo 没有SQL注入处理, 不建议引用，加入了事务
func RawUpdateUserInfo(uid int64, which, what string, tx *sql.Tx) (err error) {
	sqlStr := fmt.Sprintf("UPDATE user SET %s=? WHERE uid = ?", which)
	if tx == nil {
		_, err = dB.Exec(sqlStr, what, uid)
		return
	}
	_, err = tx.Exec(sqlStr, what, uid)
	return
}

// UpdateUserDescription 预处理解决 SQL 注入问题，加入了事务
func UpdateUserDescription(uid int64, value string, tx *sql.Tx) (err error) {
	sqlStr := "UPDATE user SET description = ? WHERE uid = ?"

	var stmt *sql.Stmt
	if tx == nil {
		stmt, err = dB.Prepare(sqlStr)
	} else {
		stmt, err = tx.Prepare(sqlStr)
	}

	if err != nil {
		return
	}
	_, err = stmt.Exec(value, uid)
	err = stmt.Close()
	if err != nil {
		utils.LoggerWarning("Statement 关闭异常!", err)
	}
	return
}

func DeleteUser(uid int64) (err error) {
	sqlStr := "DELETE FROM user WHERE uid = ?"
	_, err = dB.Exec(sqlStr, uid)
	return
}
