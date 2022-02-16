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
	discussion.Content = utils.ReplaceXSSKeywords(discussion.Content)
	discussion.Content = utils.ReplaceWildUrl(discussion.Content)
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
	content = utils.ReplaceXSSKeywords(content)
	content = utils.ReplaceWildUrl(content)
	return dao.UpdateDiscussion(id, uid, name, content)
}

func StarDiscussion(id, uid int64) (err error) {
	return dao.StarOrUnStarDiscussion(id, uid, true)
}

func UnStarDiscussion(id, uid int64) (err error) {
	return dao.StarOrUnStarDiscussion(id, uid, false)
}
