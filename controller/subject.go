package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"net/http"
)

func CtrlSubjectsGet(start, limit int, sort string, tags string) (err error, resp utils.RespData) {
	err, subjects := service.GetSubjects(start, limit, sort, tags)

	var movieTags = make([]model.MovieTag, 0)

	for _, subject := range subjects {
		var movie model.MovieTag

		movie.Name = subject.Name
		movie.Avatar = subject.Avatar
		movie.Score = subject.Score.Score
		movie.Mid = subject.Mid

		movieTags = append(movieTags, movie)
	}
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       movieTags,
	}
	return
}
