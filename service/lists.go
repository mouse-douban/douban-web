package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
	"net/http"
	"time"
)

func GetMovieList(id int64) (err error, movieList model.MovieList) {
	err, movieList = dao.SelectMovieList(id)
	if err != nil {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40013,
			Info:       "invalid request",
			Detail:     "没有这个片单",
		}, movieList
	}
	return
}

// UpdateMovieListInfo 需要对 params 做好正则检测，加入事务
func UpdateMovieListInfo(id int64, params map[string]interface{}, updateTime bool) (err error) {
	tx, err := dao.OpenTransaction()
	if err != nil {
		return
	}
	for key, value := range params {
		// 预处理防止 sql 注入
		if key == "description" || key == "name" {
			err = dao.PrepareUpdateMovieList(id, key, value, tx)
			if err != nil {
				dao.RollBackTransaction(tx)
				return
			}
			continue
		}
		err = dao.RawUpdateMovieList(id, key, value, tx)
		if err != nil {
			dao.RollBackTransaction(tx)
			return
		}
	}
	if updateTime {
		err = dao.RawUpdateMovieList(id, "date", time.Now(), tx)
		if err != nil {
			dao.RollBackTransaction(tx)
			return
		}
	}
	dao.CommitTransaction(tx)
	return
}

func CreateMovieList(uid int64, name, description string, list []int64) (err error, lid int64) {
	if name == "" {
		name = "未命名"
	}
	movieList := model.MovieList{
		Uid:         uid,
		Name:        name,
		Description: description,
		List:        list,
		Date:        time.Now(),
		Followers:   0,
	}
	return dao.InsertMovieList(movieList)
}

func DeleteMovieList(lid, uid int64) (err error) {
	err = dao.DeleteMovieList(lid, uid)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40000,
			Info:       "invalid request",
			Detail:     "删除失败",
		}
	}
	return err
}
