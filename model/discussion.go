package model

type DiscussionSnapshot struct {
	Id       int64  `json:"id,omitempty"`
	Uid      int64  `json:"uid,omitempty"`
	Mid      int64  `json:"mid,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	ReplyCnt int64  `json:"reply_cnt,omitempty"`
	Date     string `json:"date,omitempty"`
	Stars    int64  `json:"stars,omitempty"`
}
