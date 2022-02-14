package model

import "time"

// Reply 回复
type Reply struct {
	Id       int64     `json:"id"`
	Uid      int64     `json:"uid"`
	Pid      int64     `json:"pid"`  // parent_id
	Ptable   string    `json:"type"` // parent_table
	Date     time.Time `json:"date"`
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Avatar   string    `json:"avatar"`
}
