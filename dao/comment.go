package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"strings"
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
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 关闭失败, cause", err)
		}
	}(stmt)
	if err != nil {
		return
	}
	_, err = stmt.Exec(comment.Mid, comment.Uid, comment.Content, comment.Date, comment.Score, strings.Join(comment.Tag, ","), comment.Type, comment.Stars)
	return
}

func UpdateComment(id, uid int64, tag []string, content string, score int) (err error) {
	sqlStr := "UPDATE comment SET tag = ?, content = ?, score = ?, date = ? WHERE id = ? AND uid = ?"
	stmt, err := dB.Prepare(sqlStr)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 关闭失败, cause", err)
		}
	}(stmt)
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
