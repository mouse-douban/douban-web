package model

import "time"

// Comment 短评，没有快照形式
type Comment struct {
	Id       int64     `json:"-"`
	Mid      int64     `json:"mid"`
	Uid      int64     `json:"uid"`
	Tag      []string  `json:"tag"`
	Content  string    `json:"content"`
	Score    int       `json:"score"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Date     time.Time `json:"date"`
	Stars    int64     `json:"stars"`
}
