package controller

import (
	"douban-webend/dao"
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

func CtrlSubjectBaseInfoGet(mid int64) (err error, resp utils.RespData) {
	err, movie := service.GetSubjectBaseInfo(mid)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40015,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       movie,
	}
	return
}

func CtrlSubjectScopeInfoGet(mid int64, scopes []string) (err error, resp utils.RespData) {
	var info = make(map[string][]interface{})

	err = service.GetSubjectScopeInfo(mid, scopes, &info)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40015,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}

	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       info,
	}
	return
}

func CtrlSubjectCommentsGet(mid int64, start, limit int, sort, kind string) (err error, resp utils.RespData) {
	var data = make([]interface{}, 0)
	err = service.GetSubjectComments(mid, &data, start, limit, sort, kind)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       data,
	}
	return
}

func CtrlSubjectReviewsGet(mid int64, start, limit int, sort string) (err error, resp utils.RespData) {
	var data = make([]interface{}, 0)
	err = service.GetSubjectReviews(mid, &data, start, limit, sort)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       data,
	}
	return
}

func CtrlSubjectDiscussionsGet(mid int64, start, limit int, sort string) (err error, resp utils.RespData) {
	var data = make([]interface{}, 0)
	err = service.GetSubjectDiscussions(mid, &data, start, limit, sort)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       data,
	}
	return
}

func CtrlWhatSubjectsNameLike(name string) (err error, resp utils.RespData) {
	err, datum := dao.SelectSubjectNameLike(name)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       datum,
	}
	return
}
