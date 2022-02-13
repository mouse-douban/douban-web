package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"time"
)

func SelectReview(id int64) (err error, review model.Review) {
	sqlStr := "SELECT * FROM review WHERE id = ?"
	row := dB.QueryRow(sqlStr, id)
	err = row.Scan(
		&review.Id,
		&review.Name,
		&review.Mid,
		&review.Uid,
		&review.Score,
		&review.Date,
		&review.Stars,
		&review.Bads,
		&review.ReplyCnt,
		&review.Content,
	)
	return
}

func InsertReview(review model.Review) (err error) {
	sqlStr := "INSERT INTO review(name, mid, uid, score, date, content) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 未能关闭 cause", err)
		}
	}(stmt)
	_, err = stmt.Exec(review.Name, review.Mid, review.Uid, review.Score, time.Now(), review.Content)
	return
}

func DeleteReview(id, uid int64) (err error) {
	sqlStr := "DELETE FROM review WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr, id, uid)
	return
}

func UpdateReview(id, uid int64, name, content string, score int) (err error) {
	sqlStr := "UPDATE review SET name = ?, content = ?, score = ?, date = ? WHERE id = ? AND uid = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			utils.LoggerWarning("statement 未能关闭 cause", err)
		}
	}(stmt)
	_, err = stmt.Exec(name, content, score, time.Now(), id, uid)
	return
}
