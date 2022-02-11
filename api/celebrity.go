package api

import (
	"douban-webend/controller"
	"douban-webend/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func handleCelebrityGet(ctx *gin.Context) {
	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		utils.RespWithParamError(ctx, "id 参数错误")
	}
	err, resp := controller.CtrlCelebrityGet(id)
	utils.Resp(ctx, err, resp)
}
