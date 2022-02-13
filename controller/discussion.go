package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
	"time"
)

func CtrlDiscussionGet(id int64) (err error, resp utils.RespData) {
	err, discussion := service.GetDiscussion(id)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       discussion,
	}
	return
}

func CtrlDiscussionPost(uid, mid int64, name, content string) (err error, resp utils.RespData) {
	var discussion = model.Discussion{
		DiscussionSnapshot: model.DiscussionSnapshot{
			Uid:      uid,
			Mid:      mid,
			Name:     name,
			ReplyCnt: 0,
			Date:     time.Now(),
			Stars:    0,
		},
		Content: content,
	}
	err = service.CreateDiscussion(discussion)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlDiscussionDelete(id, uid int64) (err error, resp utils.RespData) {
	err = service.DeleteDiscussion(id, uid)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlDiscussionUpdate(id, uid int64, name, content string) (err error, resp utils.RespData) {
	err = service.UpdateDiscussion(id, uid, name, content)
	if err != nil {
		return
	}
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlDiscussionStar(id, uid int64, value bool) (err error, resp utils.RespData) {
	if value {
		err = service.StarDiscussion(id, uid)
	} else {
		err = service.UnStarDiscussion(id, uid)
	}
	resp = utils.NoDetailSuccessResp
	return
}
