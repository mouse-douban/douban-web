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
