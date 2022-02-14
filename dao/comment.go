package dao

import (
	"douban-webend/model"
	"douban-webend/utils"
	"strings"
	"sync"
	"time"
)

func SelectComment(id int64) (err error, comment model.Comment) {
	sqlStr := "SELECT * FROM comment WHERE id = ?"
	row := dB.QueryRow(sqlStr, id)
	var tags string
	err = row.Scan(
		&comment.Id,
		&comment.Mid,
		&comment.Uid,
		&comment.Content,
		&comment.Date,
		&comment.Score,
		&tags,
		&comment.Type,
		&comment.Stars,
	)
	comment.Tag = strings.Split(tags, ",")
	return
}

func InsertComment(comment model.Comment) (err error) {
	sqlStr := "INSERT INTO comment(mid, uid, content, date, score, tag, type, stars) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dB.Prepare(sqlStr)
	defer utils.LoggerError("Statement 关闭异常!", stmt)
	if err != nil {
		return
	}
	_, err = stmt.Exec(comment.Mid, comment.Uid, comment.Content, comment.Date, comment.Score, strings.Join(comment.Tag, ","), comment.Type, comment.Stars)
	return
}

func UpdateComment(id, uid int64, tag []string, content string, score int) (err error) {
	sqlStr := "UPDATE comment SET tag = ?, content = ?, score = ?, date = ? WHERE id = ? AND uid = ?"
	stmt, err := dB.Prepare(sqlStr)
	defer utils.LoggerError("Statement 关闭异常!", stmt)
	if err != nil {
		return
	}
	_, err = stmt.Exec(strings.Join(tag, ","), content, score, time.Now(), id, uid)
	return
}

func DeleteComment(id, uid int64) (err error) {
	sqlStr := "DELETE FROM comment WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr, id, uid)
	return
}

var mu3 = sync.Mutex{}

func StarOrUnStarComment(id, uid int64, value bool) (err error) {
	mu3.Lock()
	var v int64
	sqlStr1 := "SELECT stars FROM comment WHERE id = ? AND uid = ?"
	err = dB.QueryRow(sqlStr1, id, uid).Scan(&v)
	if err != nil {
		return
	}
	if value {
		v += 1
	} else {
		v -= 1
	}
	sqlStr2 := "UPDATE comment SET stars = ? WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr2, v, id, uid)
	mu3.Unlock()
	return
}
