package model

import "time"

type Movie struct {
	Mid    int64       `json:"mid"`
	Name   string      `json:"name"`
	Stars  int64       `json:"stars"`
	Date   time.Time   `json:"date"`
	Tags   string      `json:"tags"`
	Detail MovieDetail `json:"detail"`
	Score  MovieScore  `json:"score"`
	Plot   string      `json:"plot"`
}

type MovieScore struct {
	Total    int `json:"total"`
	TotalCnt int `json:"total_cnt"`
	Five     int `json:"five"`
	Four     int `json:"four"`
	Three    int `json:"three"`
	Two      int `json:"two"`
	One      int `json:"one"`
}

type MovieDetail struct {
	Nicknames  []string  `json:"nicknames"`
	Director   string    `json:"director"`
	Writers    []string  `json:"writers"`
	Characters []string  `json:"characters"`
	Type       []string  `json:"type"`
	Website    string    `json:"website"`
	Region     string    `json:"region"`
	Language   string    `json:"language"`
	Release    time.Time `json:"release"`
	Period     int       `json:"period"`
	IMDb       string    `json:"IMDb"`
}
