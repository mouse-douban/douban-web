package controller

import (
	"douban-webend/service"
	"douban-webend/utils"
	"net/http"
)

func CtrlSubjectsGet(start, limit int, sort string, tags string) (err error, resp utils.RespData) {
	err, subjects := service.GetSubjects(start, limit, sort, tags)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       subjects,
	}
	return
}
