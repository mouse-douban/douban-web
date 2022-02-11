package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func GetCelebrity(id int64) (err error, celebrity model.Celebrity) {
	err, celebrity = dao.SelectCelebrity(id)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40016,
			Info:       "invalid request",
			Detail:     "影人不存在",
		}
	}
	return
}
