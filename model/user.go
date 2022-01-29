package model

import (
	"golang.org/x/crypto/bcrypt"
)

type MoveList struct { // 片单
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	List []int64 `json:"list"`
}

type User struct {
	Username       string           `json:"username,omitempty"`
	Uid            int64            `json:"uid,omitempty"` // 只读
	GithubId       int64            `json:"github_id,omitempty"`
	GiteeId        int64            `json:"gitee_id,omitempty"`
	Email          string           `json:"email,omitempty"`
	Phone          string           `json:"phone,omitempty"`
	Avatar         string           `json:"avatar,omitempty"`
	Reviews        []ReviewSnapshot `json:"reviews,omitempty"`    // 可选
	MovieList      []MoveList       `json:"movie_list,omitempty"` // 可选
	Before         []Comment        `json:"before,omitempty"`     // 可选
	After          []Comment        `json:"after,omitempty"`      // 可选
	PlaintPassword string           `json:"-"`                    // 只写
}

// OAuthInfo 需要得到的基础的信息
type OAuthInfo struct {
	Username string `json:"login"`
	OAuthId  int64  `json:"id"`
	Avatar   string `json:"avatar_url"`
	PlatForm string `json:"-"`
}

func (u *User) EncryptPassword() string {
	en, _ := bcrypt.GenerateFromPassword([]byte(u.PlaintPassword), bcrypt.DefaultCost)
	return string(en)
}
