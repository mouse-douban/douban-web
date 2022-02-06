package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"encoding/json"
)

func SelectMovieList(id int64) (err error, movieList model.MovieList) {
	sqlStr := "SELECT id, uid, name, date, followers, list, description FROM movie_list WHERE id = ?"
	row := dB.QueryRow(sqlStr, id)
	var listJson string
	err = row.Scan(
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
	return
}

// RawUpdateMovieList 没有SQL注入的处理，不能更新 name, description 等字段, 加入了事务
func RawUpdateMovieList(id int64, column string, value interface{}, tx *sql.Tx) (err error) {
	sqlStr := "UPDATE movie_list SET " + column + " = ? WHERE id = ?"
	if tx == nil {
		_, err = dB.Exec(sqlStr, value, id)
	} else {
		_, err = tx.Exec(sqlStr, value, id)
	}
	return
}

func PrepareUpdateMovieList(id int64, column string, value interface{}, tx *sql.Tx) (err error) {
	sqlStr := "UPDATE movie_list SET " + column + " = ? WHERE id = ?"
	var stmt *sql.Stmt
	if tx == nil {
		stmt, err = dB.Prepare(sqlStr)
	} else {
		stmt, err = tx.Prepare(sqlStr)
	}
	if err != nil {
		return
	}
	_, err = stmt.Exec(value, id)
	if err != nil {
		return
	}
	err = stmt.Close()
	if err != nil {
		utils.LoggerWarning("关闭 Statement 失败, 原因", err)
		return
	}
	return
}

func InsertMovieList(movieList model.MovieList) (err error, id int64) {
	tx, err := OpenTransaction() // 开启一个事务
	if err != nil {
		return
	}
	// 写入不会被 SQL 注入的信息
	sqlStr := "INSERT INTO movie_list(uid, date, list) VALUES(?, ?, ?)"
	list, err := json.Marshal(movieList.List)
	if err != nil {
		RollBackTransaction(tx)
		return
	}
	res, err := tx.Exec(sqlStr, movieList.Uid, movieList.Date, string(list))
	if err != nil {
		RollBackTransaction(tx)
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		RollBackTransaction(tx)
		return
	}
	// 预处理处理 SQL 注入
	sqlStr = "UPDATE movie_list SET name = ?, description = ? WHERE id = ?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		RollBackTransaction(tx)
		return
	}
	_, err = stmt.Exec(movieList.Name, movieList.Description, id)
	if err != nil {
		RollBackTransaction(tx)
		return
	}
	err = stmt.Close()
	if err != nil {
		utils.LoggerWarning("Statement 关闭失败，原因: ", err)
		RollBackTransaction(tx)
		return
	}
	CommitTransaction(tx)
	return
}

func DeleteMovieList(lid, uid int64) (err error) {
	sqlStr := "DELETE FROM movie_list WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr, lid, uid)
	return
}
