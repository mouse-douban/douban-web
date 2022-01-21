package model

type Comment struct {
	Id       int64    `json:"-"`
	Tag      []string `json:"tag"`
	Content  string   `json:"content"`
	Score    int64    `json:"score"`
	Username string   `json:"username"`
	Uid      int64    `json:"uid"`
	Type     string   `json:"type"`
	Date     string   `json:"date"`
	Stars    int64    `json:"stars"`
}
