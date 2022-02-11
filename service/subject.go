package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func GetSubjects(start, limit int, sort string, tags string) (err error, subjects []model.Movie) {
	err, subjects = dao.SelectSubjects(tags, orderBys[sort])
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	end := start + limit
	if start+limit > len(subjects) {
		end = len(subjects)
	}
	return nil, subjects[start:end]
}

func GetSubjectBaseInfo(mid int64) (err error, movie model.Movie) {
	err, movie = dao.SelectSubjectBaseInfo(mid)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40015,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
	}
	return err, movie
}

func GetSubjectScopeInfo(mid int64, scopes []string, info *map[string][]interface{}) (err error) {
	for _, scope := range scopes {
		switch scope {
		case "comments":
			var in = make([]interface{}, 0)
			err = GetSubjectComments(mid, &in, 0, 6, "hotest", "after")
			if err != nil {
				return
			}
			(*info)["comments"] = in
		case "reviews":
			var in = make([]interface{}, 0)
			err = GetSubjectReviews(mid, &in, 0, 6, "hotest")
			if err != nil {
				return
			}
			(*info)["reviews"] = in
		case "discussions":
			var in = make([]interface{}, 0)
			err = GetSubjectDiscussions(mid, &in, 0, 6, "hotest")
			if err != nil {
				return
			}
			(*info)["discussions"] = in
		}
	}
	return
}

func GetSubjectComments(mid int64, comments *[]interface{}, start, limit int, sort, kind string) (err error) {
	err = dao.SelectSubjectComments(mid, orderBys[sort], kind, comments)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	end := start + limit
	if end > len(*comments) {
		end = len(*comments)
	}
	*comments = (*comments)[start:end]
	return
}

func GetSubjectReviews(mid int64, reviews *[]interface{}, start, limit int, sort string) (err error) {
	err = dao.SelectSubjectReviews(mid, orderBys[sort], reviews)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	end := start + limit
	if end > len(*reviews) {
		end = len(*reviews)
	}
	*reviews = (*reviews)[start:end]
	return
}

func GetSubjectDiscussions(mid int64, discussions *[]interface{}, start, limit int, sort string) (err error) {
	err = dao.SelectSubjectDiscussions(mid, orderBys[sort], discussions)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	end := start + limit
	if end > len(*discussions) {
		end = len(*discussions)
	}
	*discussions = (*discussions)[start:end]
	return
}
