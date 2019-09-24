package api

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResp struct {
	Code int64                           `json:"code"` // 错误码
	Msg  string                          `json:"msg"`  // 错误详
	Data *models.GetEditDistanceResqData `json:"data"` // 返回的数据
}

func HandleError(c *gin.Context, err ApiResp) {
	resp := &ApiResp{
		Code: err.Code,
		Msg:  err.Msg,
		Data: err.Data,
	}
	c.JSON(http.StatusOK, resp)
}
