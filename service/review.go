package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
)

func GetReview(id int64) (err error, review model.Review) {
	err, review = dao.SelectReview(id)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40020,
			Info:       "invalid request",
			Detail:     "影评不存在",
		}
	}
	return
}

func CreateReview(review model.Review) (err error) {
	err = dao.InsertReview(review)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40018,
			Info:       "invalid request",
			Detail:     "创建失败",
		}
		return
	}
	err = UpdateSubjectScore(review.Mid, review.Score)
	if err != nil {
		return
	}
	return
}

func DeleteReview(id, uid int64) (err error) {
	return dao.DeleteReview(id, uid)
}

func UpdateReview(id, uid int64, name, content string, score int) (err error) {
	return dao.UpdateReview(id, uid, name, content, score)
}

func StarReview(id, uid int64) (err error) {
	return dao.StarOrUnStarReview(id, uid, true)
}

func UnStarReview(id, uid int64) (err error) {
	return dao.StarOrUnStarReview(id, uid, false)
}

func BadReview(id, uid int64) (err error) {
	return dao.BadOrUnBadReview(id, uid, true)
}

func UnBadReview(id, uid int64) (err error) {
	return dao.BadOrUnBadReview(id, uid, false)
}
