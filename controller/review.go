package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
)

func CtrlReviewGet(id int64) (err error, resp utils.RespData) {
	err, review := service.GetReview(id)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       review,
	}
	return
}

func CtrlReviewPost(mid, uid int64, name, content string, score int) (err error, resp utils.RespData) {
	var review = model.Review{
		ReviewSnapshot: model.ReviewSnapshot{
			Mid:   mid,
			Uid:   uid,
			Name:  name,
			Score: score,
		},
		Content: content,
	}
	err = service.CreateReview(review)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlReviewDelete(id, uid int64) (err error, resp utils.RespData) {
	err = service.DeleteReview(id, uid)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlReviewUpdate(id, uid int64, name, content string, score int) (err error, resp utils.RespData) {
	err = service.UpdateReview(id, uid, name, content, score)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}
