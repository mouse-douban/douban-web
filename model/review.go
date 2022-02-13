package model

import "time"

type ReviewSnapshot struct {
	Id       int64     `json:"id"`
	Mid      int64     `json:"mid"`
	Uid      int64     `json:"uid"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	Score    int       `json:"score"`
	Date     time.Time `json:"date"`
	Stars    int64     `json:"stars"`
	Bads     int64     `json:"bads"`
	ReplyCnt int64     `json:"reply_cnt"`
	Brief    string    `json:"brief"`
}

type Review struct {
	ReviewSnapshot
	Content string `json:"content"`
}
