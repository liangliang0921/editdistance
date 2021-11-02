package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"myProject/editdistance/models"
	"net/http"
)

type GetEditDistanceReq struct {
	NewString string `json:"newstring"` // 新字符串
	OldString string `json:"oldstring"` // 旧字符串
}

type GetEditDistanceResq struct {
	Code int                             `json:"code"` // 错误码
	Msg  string                          `json:"msg"`  // 错误详情
	Data *models.GetEditDistanceResqData `json:"data"` // 返回的数据
}

func GetShortestEditDistance(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		HandleError(c, NewError(ErrorGetRawFailed))
		fmt.Printf("GetRawData failed|msg=%v", err)
		return
	}
	req := GetEditDistanceReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		HandleError(c, NewError(ErrorUnmarshFailed))
		fmt.Printf("Unmarshal failed|msg=%v", err)
		return
	}
	editDistanceData := models.GetEditDistance(req.NewString, req.OldString)
	c.JSON(http.StatusOK, GetEditDistanceResq{
		Code: 0,
		Msg:  "success",
		Data: editDistanceData,
	})
}
