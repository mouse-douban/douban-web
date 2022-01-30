package controller

import (
	"douban-webend/service"
	"douban-webend/utils"
	"encoding/json"
	"net/http"
)

func CtrlMovieListGet(lid int64) (err error, resp utils.RespData) {
	err, list := service.GetMovieList(lid)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       list,
	}
	return
}

func CtrlMovieListCreate(uid int64, name, description string, list []int64) (err error, resp utils.RespData) {
	err, lid := service.CreateMovieList(uid, name, description, list)
	if err != nil {
		return
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data: struct {
			Id int64 `json:"id,omitempty"`
		}{
			Id: lid,
		},
	}
	return
}

func CtrlDeleteMovieList(lid, uid int64) (err error, resp utils.RespData) {
	err = service.DeleteMovieList(lid, uid)
	return err, utils.NoDetailSuccessResp
}

func CtrlUpdateMovieList(lid, uid int64, params map[string]interface{}) (err error, resp utils.RespData) {
	err, list := service.GetMovieList(lid)
	if err != nil {
		return
	}
	if list.Uid != uid {
		err = utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40014,
			Info:       "invalid request",
			Detail:     "不能修改别人的片单",
		}
		return
	}
	err = service.UpdateMovieListInfo(lid, params, true)
	return err, utils.NoDetailSuccessResp
}

func CtrlMovieListMovieAdd(lid, uid int64, newMids []int64) (err error, resp utils.RespData) {
	err, list := service.GetMovieList(lid)
	if err != nil {
		return
	}
	if list.Uid != uid {
		err = utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40014,
			Info:       "invalid request",
			Detail:     "不能修改别人的片单",
		}
		return
	}
	flag := make(map[int64]bool)
	for _, mid := range list.List {
		flag[mid] = true
	}
	for _, mid := range newMids {
		if check, ok := flag[mid]; check && ok {
			continue
		}
		list.List = append(list.List, mid)
		flag[mid] = true
	}
	bytes, err := json.Marshal(list.List)
	if err != nil {
		return
	}
	err = service.UpdateMovieListInfo(lid, map[string]interface{}{"list": string(bytes)}, true)
	resp = utils.NoDetailSuccessResp
	return
}

func CtrlMovieListMovieRemove(lid, uid int64, removeMids []int64) (err error, resp utils.RespData) {
	err, list := service.GetMovieList(lid)
	if err != nil {
		return
	}
	if list.Uid != uid {
		err = utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40014,
			Info:       "invalid request",
			Detail:     "不能修改别人的片单",
		}
		return
	}

	newMids := make([]int64, 0)
	flag := make(map[int64]bool)

	for _, mid := range removeMids {
		flag[mid] = true
	}
	for _, mid := range list.List {
		if _, ok := flag[mid]; !ok {
			newMids = append(newMids, mid)
		}
	}

	bytes, err := json.Marshal(newMids)
	if err != nil {
		return
	}
	err = service.UpdateMovieListInfo(lid, map[string]interface{}{"list": string(bytes)}, true)
	resp = utils.NoDetailSuccessResp
	return
}
