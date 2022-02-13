package model

import "time"

type DiscussionSnapshot struct {
	Id       int64     `json:"id"`
	Uid      int64     `json:"uid"`
	Mid      int64     `json:"mid"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	ReplyCnt int64     `json:"reply_cnt"`
	Date     time.Time `json:"date"`
	Avatar   string    `json:"avatar"`
	Stars    int64     `json:"stars"`
}

type Discussion struct {
	DiscussionSnapshot
	Content string `json:"content"`
}
