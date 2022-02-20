package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"sync"
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

// 唆一锁
var mu1 = &sync.Mutex{}
var mu2 = &sync.Mutex{}
var mu4 = &sync.Mutex{}

func IncreaseReviewReplyCnt(id int64) (err error) {
	mu4.Lock()
	defer mu4.Unlock()
	var cnt int64
	row := dB.QueryRow("SELECT reply_cnt FROM review WHERE id = ?", id)
	err = row.Scan(&cnt)
	if err != nil {
		return
	}
	cnt++
	_, err = dB.Exec("UPDATE review SET reply_cnt = ? WHERE id = ?", cnt, id)
	return
}

func StarOrUnStarReview(id, uid int64, value bool) (err error) {
	return starOrBadReview(id, uid, "stars", mu1, value)
}

func BadOrUnBadReview(id, uid int64, value bool) (err error) {
	return starOrBadReview(id, uid, "bads", mu2, value)
}

func starOrBadReview(id, uid int64, kind string, mu *sync.Mutex, value bool) (err error) {
	mu.Lock()
	defer mu.Unlock()
	var v int64
	sqlStr1 := "SELECT " + kind + " FROM review WHERE id = ? AND uid = ?"
	err = dB.QueryRow(sqlStr1, id, uid).Scan(&v)
	if err != nil {
		return
	}
	if value {
		v += 1
	} else {
		v -= 1
	}
	sqlStr2 := "UPDATE review SET " + kind + " = ? WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr2, v, id, uid)
	return
}
