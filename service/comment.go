package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func GetComment(id int64) (err error, comment model.Comment) {
	err, comment = dao.SelectComment(id)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40017,
			Info:       "invalid request",
			Detail:     "短评不存在",
		}
	}
	return
}

func CreateComment(comment model.Comment) (err error) {
	err = dao.InsertComment(comment)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40018,
			Info:       "invalid request",
			Detail:     "创建失败",
		}
		return
	}
	err = UpdateSubjectScore(comment.Mid, comment.Score)
	if err != nil {
		return
	}
	return err
}

func UpdateComment(id, uid int64, tag []string, content string, score int) (err error) {
	return dao.UpdateComment(id, uid, tag, content, score)
}

func DeleteComment(id, uid int64) (err error) {
	return dao.DeleteComment(id, uid)
}
