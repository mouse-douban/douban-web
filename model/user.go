package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type MovieList struct { // 片单
	Id          int64     `json:"id"`
	Uid         int64     `json:"uid"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Followers   int64     `json:"followers"`
	List        []int64   `json:"list"`
	Description string    `json:"description"`
}

type Follow struct {
	Users map[int64]bool `json:"users"`
	Lists map[int64]bool `json:"movie_lists"`
}

type User struct {
	Username       string           `json:"username"`
	Uid            int64            `json:"uid"` // 只读
	GithubId       int64            `json:"github_id"`
	GiteeId        int64            `json:"gitee_id"`
	Email          string           `json:"email"`
	Phone          string           `json:"phone"`
	Avatar         string           `json:"avatar"`
	Description    string           `json:"description"`
	Following      Follow           `json:"following"`
	Reviews        []ReviewSnapshot `json:"reviews,omitempty"`    // 可选
	MovieList      []MovieList      `json:"movie_list,omitempty"` // 可选
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
