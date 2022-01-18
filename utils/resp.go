package utils

// resp 返回工具类

import (
	"encoding/json"
	"fmt"
)

type ServerError struct {
	HttpStatus int    `json:"-"`
	Status     int    `json:"status"`
	Info       string `json:"info"`
	Detail     string `json:"detail"`
}

type Detail struct {
	Detail string `json:"detail"`
}

type RespDetail struct {
	HttpStatus int    `json:"-"`
	Status     int    `json:"status"`
	Info       string `json:"info"`
	Data       Detail `json:"data"`
}

var (
	ServerInternalError = ServerError{ // 服务器内部错误
		HttpStatus: 500,
		Status:     50000,
		Info:       "server error",
		Detail:     "no detail message",
	}
	QueryParamError = ServerError{ // 请求参数错误，没有 Detail，需要 Copy 后补充
		HttpStatus: 422,
		Status:     42200,
		Info:       "invalid params",
	}
	NoDetailSuccessResp = RespDetail{
		HttpStatus: 200,
		Status:     22222,
		Info:       InfoSuccess,
		Data: Detail{
			Detail: "no detail message",
		},
	}

	ServerInternalErrorJSON = "{\n  \"status\": 50000,\n  \"info\": \"server error\",\n  \"data\": {\n    \"detail\": \"no detail message\"\n  }\n}"
	InfoSuccess             = "success"
)

func (s ServerError) Error() string {
	return fmt.Sprintf("Status: %d Info: %s Detail: %s", s.Status, s.Info, s.Detail)
}

func (s ServerError) Copy() ServerError {
	return ServerError{
		HttpStatus: s.HttpStatus,
		Status:     s.Status,
		Detail:     s.Detail,
		Info:       s.Info,
	}
}

func (r RespDetail) RespJSON() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ServerInternalErrorJSON
	}
	return string(marshal)
}

func (s ServerError) RespDetail() RespDetail {
	var resp RespDetail
	resp.Status = s.Status
	resp.Info = s.Info
	resp.Data.Detail = s.Detail
	return resp
}
