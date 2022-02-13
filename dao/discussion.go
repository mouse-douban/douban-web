package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"sync"
	"time"
)

func SelectDiscussion(id int64) (err error, discussion model.Discussion) {
	sqlStr := "SELECT * FROM discussion WHERE id = ?"
	row := dB.QueryRow(sqlStr, id)
	err = row.Scan(
		&discussion.Id,
		&discussion.Uid,
		&discussion.Mid,
		&discussion.Name,
		&discussion.ReplyCnt,
		&discussion.Date,
		&discussion.Stars,
		&discussion.Content,
	)
	return
}

func InsertDiscussion(discussion model.Discussion) (err error) {
	sqlStr := "INSERT INTO discussion(uid, mid, name, reply_cnt, date, stars, content) VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 未关闭, cause", err)
		}
	}(stmt)
	_, err = stmt.Exec(discussion.Uid, discussion.Mid, discussion.Name, discussion.ReplyCnt, discussion.Date, discussion.Stars, discussion.Content)
	return
}

func DeleteDiscussion(id, uid int64) (err error) {
	sqlStr := "DELETE FROM discussion WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr, id, uid)
	return
}

func UpdateDiscussion(id, uid int64, name, content string) (err error) {
	sqlStr := "UPDATE discussion SET name = ?, content = ?, date = ? WHERE id = ? AND uid = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 未关闭, cause", err)
		}
	}(stmt)
	_, err = stmt.Exec(name, content, time.Now(), id, uid)
	return
}

var mu = sync.Mutex{}

func StarOrUnStarDiscussion(id, uid int64, value bool) (err error) {
	mu.Lock()
	var v int64
	sqlStr1 := "SELECT stars FROM discussion WHERE id = ? AND uid = ?"
	err = dB.QueryRow(sqlStr1, id, uid).Scan(&v)
	if err != nil {
		return
	}
	if value {
		v += 1
	} else {
		v -= 1
	}
	sqlStr2 := "UPDATE discussion SET stars = ? WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr2, v, id, uid)
	mu.Unlock()
	return
}
