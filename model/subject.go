package model

import "time"

type MovieTag struct { // 用于返回电影列表
	Name   string `json:"name"`
	Score  string `json:"score"` // 折算成10.0分制
	Avatar string `json:"avatar"`
	Mid    int64  `json:"mid"`
}

type Movie struct {
	Mid         int64       `json:"mid"`
	Name        string      `json:"name"`
	Stars       int64       `json:"stars"` // 平均星数
	Date        time.Time   `json:"date"`
	Tags        string      `json:"tags"`
	Avatar      string      `json:"avatar"`
	Detail      MovieDetail `json:"detail"`
	Score       MovieScore  `json:"score"`
	Plot        string      `json:"plot"`
	Celebrities []int64     `json:"celebrities"` // 所有演职员 id
}

type MovieScore struct {
	Score    string `json:"score"`     // 折算成10.0分制
	TotalCnt int    `json:"total_cnt"` // 总评分人数
	Five     string `json:"five"`      // 五星占比
	Four     string `json:"four"`
	Three    string `json:"three"`
	Two      string `json:"two"`
	One      string `json:"one"`
}

type MovieDetail struct {
	Nicknames  []string `json:"nicknames"`
	Director   string   `json:"director"`
	Writers    []string `json:"writers"`
	Characters []string `json:"characters"`
	Type       []string `json:"type"`
	Website    string   `json:"website"`
	Region     string   `json:"region"`
	Language   string   `json:"language"`
	Release    string   `json:"release"`
	Period     int      `json:"period"`
	IMDb       string   `json:"IMDb"`
}
