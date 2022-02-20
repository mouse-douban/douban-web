package service

import (
	"container/list"
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
	"time"
)

func GetSubjectReplies(pid int64, ptable string, start, limit int) (err error, replies []model.Reply) {
	err, replies = dao.SelectRepliesFromPidAndPtable(pid, ptable, start, limit, false)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40000,
			Info:       "invalid request",
			Detail:     "记录不存在",
		}
	}
	return err, replies
}

func GetAllRepliesOfAReply(pid int64) (err error, replies []model.Reply) {
	// BFS
	replies = make([]model.Reply, 0)
	err, reply := dao.SelectReply(pid)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40000,
			Info:       "invalid request",
			Detail:     "记录不存在",
		}
		return
	}
	queue := list.New()
	queue.PushFront(reply)
	for queue.Front() != nil {
		front := queue.Front()
		parent := front.Value.(model.Reply)
		err, children := dao.SelectRepliesFromPidAndPtable(parent.Id, "reply", -1, -1, true)
		if err != nil {
			err = utils.ServerError{
				HttpStatus: 400,
				Status:     40000,
				Info:       "invalid request",
				Detail:     "记录不存在",
			}
			return err, nil
		}
		for _, child := range children {
			queue.PushBack(child)
		}
		queue.Remove(front)
		replies = append(replies, parent)
	}
	replies = replies[1:]
	return
}

func CreateReply(reply model.Reply) (err error) {
	if reply.Ptable == "reply" {
		err, parent := dao.SelectReply(reply.Pid)
		if err != nil {
			return utils.ServerError{
				HttpStatus: 400,
				Status:     40000,
				Info:       "invalid request",
				Detail:     "记录不存在",
			}
		}
		// 自动加上这个
		reply.Content = "回复 " + parent.Username + " : " + reply.Content
	}
	switch reply.Ptable {
	// 回复的回复不算回复(雾
	case "review":
		err = dao.IncreaseReviewReplyCnt(reply.Pid)
	case "discussion":
		err = dao.IncreaseDiscussionReplyCnt(reply.Pid)
	}
	if err != nil {
		return utils.ServerError{
			HttpStatus: 400,
			Status:     40000,
			Info:       "invalid request",
			Detail:     "记录不存在",
		}
	}
	reply.Date = time.Now()
	reply.Content = utils.ReplaceXSSKeywords(reply.Content)
	reply.Content = utils.ReplaceWildUrl(reply.Content)
	return dao.InsertReply(reply)
}

func DeleteReply(id, uid int64) (err error) {
	// 删掉所有子 reply
	err, replies := GetAllRepliesOfAReply(id)
	if err != nil {
		return
	}
	for _, reply := range replies {
		err = dao.DeleteReply(reply.Id, reply.Uid)
		if err != nil {
			return
		}
	}
	return dao.DeleteReply(id, uid)
}
