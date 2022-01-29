package model

import "time"

// Comment 短评，没有快照形式
type Comment struct {
	Id       int64     `json:"-"`
	Tag      []string  `json:"tag"`
	Content  string    `json:"content"`
	Score    int64     `json:"score"`
	Username string    `json:"username"`
	Uid      int64     `json:"uid"`
	Type     string    `json:"type"`
	Date     time.Time `json:"date"`
	Stars    int64     `json:"stars"`
}
