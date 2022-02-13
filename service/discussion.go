package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func GetDiscussion(id int64) (err error, discussion model.Discussion) {
	err, discussion = dao.SelectDiscussion(id)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40019,
			Info:       "invalid request",
			Detail:     "讨论不存在",
		}
	}
	return
}

func CreateDiscussion(discussion model.Discussion) (err error) {
	err = dao.InsertDiscussion(discussion)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40018,
			Info:       "invalid request",
			Detail:     "创建失败",
		}
	}
	return
}

func DeleteDiscussion(id, uid int64) (err error) {
	return dao.DeleteDiscussion(id, uid)
}

func UpdateDiscussion(id, uid int64, name, content string) (err error) {
	return dao.UpdateDiscussion(id, uid, name, content)
}
