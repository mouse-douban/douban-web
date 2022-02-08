package model

import "time"

type MovieTag struct { // 用于返回电影列表
	Name   string `json:"name,omitempty"`
	Score  string `json:"score,omitempty"` // 折算成10.0分制
	Avatar string `json:"avatar,omitempty"`
	Mid    int64  `json:"mid,omitempty"`
}

type Movie struct {
	Mid         int64       `json:"mid,omitempty"`
	Name        string      `json:"name,omitempty"`
	Stars       int64       `json:"stars,omitempty"` // 平均星数
	Date        time.Time   `json:"date,omitempty"`
	Tags        string      `json:"tags,omitempty"`
	Avatar      string      `json:"avatar,omitempty"`
	Detail      MovieDetail `json:"detail,omitempty"`
	Score       MovieScore  `json:"score,omitempty"`
	Plot        string      `json:"plot,omitempty"`
	Celebrities []int64     `json:"celebrities,omitempty"` // 所有演职员 id
}

type MovieScore struct {
	Score    string `json:"score,omitempty"`     // 折算成10.0分制
	TotalCnt int    `json:"total_cnt,omitempty"` // 总评分人数
	Five     string `json:"five,omitempty"`      // 五星占比
	Four     string `json:"four,omitempty"`
	Three    string `json:"three,omitempty"`
	Two      string `json:"two,omitempty"`
	One      string `json:"one,omitempty"`
}

type MovieDetail struct {
	Nicknames  []string  `json:"nicknames,omitempty"`
	Director   string    `json:"director,omitempty"`
	Writers    []string  `json:"writers,omitempty"`
	Characters []string  `json:"characters,omitempty"`
	Type       []string  `json:"type,omitempty"`
	Website    string    `json:"website,omitempty"`
	Region     string    `json:"region,omitempty"`
	Language   string    `json:"language,omitempty"`
	Release    time.Time `json:"release,omitempty"`
	Period     int       `json:"period,omitempty"`
	IMDb       string    `json:"IMDb,omitempty"`
}
