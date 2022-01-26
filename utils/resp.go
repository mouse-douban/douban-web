package utils

// resp 返回工具类

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

type RespData struct {
	HttpStatus int         `json:"-"`
	Status     int         `json:"status"`
	Info       string      `json:"info"`
	Data       interface{} `json:"data"` // Data 结构体信息
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
	NoDetailSuccessResp = RespData{
		HttpStatus: 200,
		Status:     20220,
		Info:       InfoSuccess,
		Data: Detail{
			Detail: "no detail message",
		},
	}

	InfoSuccess = "success"
)

func (s ServerError) Error() string {
	return fmt.Sprintf("Status: %d Info: %s Detail: %s", s.Status, s.Info, s.Detail)
}

func (s ServerError) CopyWithNewDetail(detail string) ServerError {
	return ServerError{
		HttpStatus: s.HttpStatus,
		Status:     s.Status,
		Detail:     detail,
		Info:       s.Info,
	}
}

func (s ServerError) Copy() ServerError {
	return ServerError{
		HttpStatus: s.HttpStatus,
		Status:     s.Status,
		Detail:     s.Detail,
		Info:       s.Info,
	}
}

func (s ServerError) GetDetail() RespDetail {
	var resp RespDetail
	resp.Status = s.Status
	resp.Info = s.Info
	resp.Data.Detail = s.Detail
	return resp
}

func Resp(ctx *gin.Context, err error, data RespData) {
	if err != nil {
		RespWithError(ctx, err)
		return
	}
	RespWithData(ctx, data)
}

func Abort(ctx *gin.Context, err error, data RespData) {
	if err != nil {
		AbortWithError(ctx, err)
		return
	}
	RespWithData(ctx, data)
	ctx.Abort()
}

func RespWithError(ctx *gin.Context, error error) {
	e, ok := error.(ServerError)
	if ok {
		ctx.JSON(e.HttpStatus, e.GetDetail())
	} else {
		e = ServerInternalError
		ctx.JSON(e.HttpStatus, e.GetDetail())
	}
}

func RespWithDetail(ctx *gin.Context, detail RespDetail) {
	ctx.JSON(detail.HttpStatus, detail)
}

func RespWithData(ctx *gin.Context, data RespData) {
	ctx.JSON(data.HttpStatus, data)
}

func AbortWithError(ctx *gin.Context, error error) {
	RespWithError(ctx, error)
	ctx.Abort()
}

func AbortWithDetail(ctx *gin.Context, detail RespDetail) {
	ctx.JSON(detail.HttpStatus, detail)
	ctx.Abort()
}

func AbortWithInternalError(ctx *gin.Context) {
	ctx.JSON(ServerInternalError.HttpStatus, ServerInternalError.GetDetail())
	ctx.Abort()
}

func SuccessWithNoContent(ctx *gin.Context) {
	ctx.JSON(NoDetailSuccessResp.HttpStatus, NoDetailSuccessResp)
}

func AbortWithParamError(ctx *gin.Context, detail string) {
	e := QueryParamError.Copy()
	e.Detail = detail
	RespWithError(ctx, e)
	ctx.Abort()
}

func RespWithParamError(ctx *gin.Context, detail string) {
	e := QueryParamError.Copy()
	e.Detail = detail
	RespWithError(ctx, e)
}
