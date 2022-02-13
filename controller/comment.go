package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"time"
)

func CtrlCommentGet(id int64) (err error, resp utils.RespData) {
	err, comment := service.GetComment(id)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       comment,
	}
	return
}

func CtrlCommentPost(mid, uid int64, content string, date time.Time, score int, tag []string, kind string, stars int64) (err error, resp utils.RespData) {
	var comment = model.Comment{
		Mid:     mid,
		Uid:     uid,
		Tag:     tag,
		Content: content,
		Score:   score,
		Type:    kind,
		Date:    date,
		Stars:   stars,
	}
	err = service.CreateComment(comment)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlCommentDelete(id, uid int64) (err error, resp utils.RespData) {
	err = service.DeleteComment(id, uid)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlCommentUpdate(id, uid int64, tag []string, content string, score int) (err error, resp utils.RespData) {
	err = service.UpdateComment(id, uid, tag, content, score)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}
