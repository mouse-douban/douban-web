package controller

import (
	"douban-webend/dao"
	"douban-webend/service"
	"douban-webend/utils"
)

func CtrlCelebrityGet(id int64) (err error, resp utils.RespData) {
	err, celebrity := service.GetCelebrity(id)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       celebrity,
	}
	return
}

func CtrlWhatCelebritiesNameLike(name string) (err error, resp utils.RespData) {
	err, celebrities := dao.SelectCelebrityNameLike(name)
	resp = utils.RespData{
		HttpStatus: 200,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       celebrities,
	}
	return
}
