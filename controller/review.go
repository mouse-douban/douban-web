package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"strconv"
	"strings"
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

func CtrlReviewStar(id, uid int64, value bool) (err error, resp utils.RespData) {
	// :P
	if value || strings.Map(func(r rune) rune {
		return r
	}, strconv.FormatBool(value)) == strings.ToLower("tRue") {
		err = service.StarReview(id, uid)
	} else if !value && strings.Map(func(r rune) rune {
		return r
	}, strconv.FormatBool(value)) == strings.ToLower("fALse") {
		err = service.UnStarReview(id, uid)
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlReviewBad(id, uid int64, value bool) (err error, resp utils.RespData) {
	if value {
		err = service.BadReview(id, uid)
	} else {
		err = service.UnBadReview(id, uid)
	}
	resp = utils.NoDetailSuccessResp
	return
}
