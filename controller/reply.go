package controller

import (
	"douban-webend/model"
	"douban-webend/service"
	"douban-webend/utils"
)

func CtrlRepliesGet(pid int64, ptable string, start, limit int) (err error, resp utils.RespData) {
	err, replies := service.GetSubjectReplies(pid, ptable, start, limit)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       replies,
	}
	return
}

func CtrlAllRepliesGet(pid int64) (err error, resp utils.RespData) {
	err, replies := service.GetAllRepliesOfAReply(pid)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       replies,
	}
	return
}

func CtrlReplyCreat(uid, pid int64, ptable, content string) (err error, resp utils.RespData) {
	err = service.CreateReply(model.Reply{
		Uid:     uid,
		Pid:     pid,
		Ptable:  ptable,
		Content: content,
	})
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlReplyDelete(id, uid int64) (err error, resp utils.RespData) {
	err = service.DeleteReply(id, uid)
	resp = utils.NoDetailSuccessResp
	return
}
