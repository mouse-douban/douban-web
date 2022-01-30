package model

type DiscussionSnapshot struct {
	Id       int64  `json:"id"`
	Uid      int64  `json:"uid"`
	Mid      int64  `json:"mid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	ReplyCnt int64  `json:"reply_cnt"`
	Date     string `json:"date"`
}
