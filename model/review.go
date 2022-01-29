package model

import "time"

type ReviewSnapshot struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Uid      int64     `json:"uid"`
	Avatar   string    `json:"avatar"`
	Score    int64     `json:"score"`
	Date     time.Time `json:"date"`
	Stars    int64     `json:"stars"`
	Bads     int64     `json:"bads"`
	ReplyCnt int64     `json:"reply_cnt"`
	Brief    string    `json:"brief"`
}
