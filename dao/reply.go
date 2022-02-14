package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"fmt"
)

func Do() {
	sqlStr := "SELECT r.id, r.uid, r.pid, r.ptable, r.date, u.username, r.content, u.avatar FROM reply r JOIN user u ON u.uid = r.uid AND r.pid = ? AND r.ptable = ? ORDER BY date DESC LIMIT ? OFFSET ?"
	rows, err := dB.Query(sqlStr, 4, "discussion", 20, 0)
	fmt.Println(rows.Next(), err)
}

// SelectRepliesFromPidAndPtable
// pid -> parent id
// ptable -> 取 discussion, review, comment, reply
func SelectRepliesFromPidAndPtable(pid int64, ptable string, start, limit int, noLimit bool) (err error, replies []model.Reply) {
	replies = make([]model.Reply, 0)
	var sqlStr string
	var rows *sql.Rows
	if noLimit {
		sqlStr = "SELECT r.id, r.uid, r.pid, r.ptable, r.date, u.username, r.content, u.avatar FROM reply r JOIN user u ON u.uid = r.uid AND r.pid = ? AND r.ptable = ? ORDER BY date DESC"
		rows, err = dB.Query(sqlStr, pid, ptable)
	} else {
		sqlStr = "SELECT r.id, r.uid, r.pid, r.ptable, r.date, u.username, r.content, u.avatar FROM reply r JOIN user u ON u.uid = r.uid AND r.pid = ? AND r.ptable = ? ORDER BY date DESC LIMIT ? OFFSET ?"
		rows, err = dB.Query(sqlStr, pid, ptable, limit, start)

	}
	if err != nil {
		return
	}

	defer utils.LoggerError("rows 关闭异常!", rows)

	for rows.Next() {
		var reply model.Reply
		err = rows.Scan(
			&reply.Id,
			&reply.Uid,
			&reply.Pid,
			&reply.Ptable,
			&reply.Date,
			&reply.Username,
			&reply.Content,
			&reply.Avatar,
		)
		if err != nil {
			return
		}
		replies = append(replies, reply)
	}
	return
}

func SelectReply(id int64) (err error, reply model.Reply) {
	sqlStr := "SELECT r.id, r.uid, r.pid, r.ptable, r.date, u.username, r.content, u.avatar FROM reply r JOIN user u ON u.uid = r.uid AND r.id = ?"
	row := dB.QueryRow(sqlStr, id)
	err = row.Scan(
		&reply.Id,
		&reply.Uid,
		&reply.Pid,
		&reply.Ptable,
		&reply.Date,
		&reply.Username,
		&reply.Content,
		&reply.Avatar,
	)
	return
}

func InsertReply(reply model.Reply) (err error) {
	sqlStr := "INSERT INTO reply(uid, pid, ptable, date, content) VALUES(?, ?, ?, ?, ?)"
	stmt, err := dB.Prepare(sqlStr)
	defer utils.LoggerError("statement 关闭失败", stmt)
	_, err = stmt.Exec(reply.Uid, reply.Pid, reply.Ptable, reply.Date, reply.Content)
	return
}

func DeleteReply(id int64, uid int64) (err error) {
	sqlStr := "DELETE FROM reply WHERE id = ? AND uid = ?"
	_, err = dB.Exec(sqlStr, id, uid)
	return
}
