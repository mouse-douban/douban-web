package model

type DiscussionSnapshot struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Uid      int    `json:"uid"`
	ReplyCnt int    `json:"reply_cnt"`
	Date     string `json:"date"`
}
