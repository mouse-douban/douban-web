package service

import (
	"douban-webend/dao"
	"douban-webend/model"
)

func GetSubjects(start, limit int, sort string, tags string) (err error, subjects []model.Movie) {
	err, subjects = dao.SelectSubjects(tags, oderBys[sort])
	if err != nil {
		return
	}
	end := start + limit
	if start+limit > len(subjects) {
		end = len(subjects)
	}
	return nil, subjects[start:end]
}

func GetSubjectBaseInfo(mid int64) (err error, movie model.Movie) {
	return dao.SelectSubjectBaseInfo(mid)
}

func GetSubjectScopeInfo(mid int64, scopes []string, info []interface{}) (err error) {
	for _, scope := range scopes {
		err = dao.SelectSubjectScopeInfo(mid, scope, info)
		if err != nil {
			return
		}
	}
	return
}
