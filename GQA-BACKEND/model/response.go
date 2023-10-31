package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Result(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

func ResponseSuccessMessage(message string, c *gin.Context) {
	Result(global.GqaConfig.System.SuccessCode, nil, message, c)
}

func ResponseSuccessMessageWithLog(message string, c *gin.Context) {
	global.GqaSLogger.Info(message)
	ResponseSuccessMessage(message, c)
}

func ResponseSuccessData(data interface{}, c *gin.Context) {
	Result(global.GqaConfig.System.SuccessCode, data, global.GqaConfig.System.SuccessMessage, c)
}

func ResponseSuccessMessageData(data interface{}, message string, c *gin.Context) {
	Result(global.GqaConfig.System.SuccessCode, data, message, c)
}

func ResponseSuccessMessageDataWithLog(data interface{}, message string, c *gin.Context) {
	global.GqaSLogger.Info(message)
	ResponseSuccessMessageData(data, message, c)
}

func ResponseErrorMessage(message string, c *gin.Context) {
	Result(global.GqaConfig.System.ErrorCode, nil, message, c)
}

func ResponseErrorMessageWithLog(message string, c *gin.Context) {
	global.GqaSLogger.Error(message)
	ResponseErrorMessage(message, c)
}

func ResponseErrorData(data interface{}, c *gin.Context) {
	Result(global.GqaConfig.System.ErrorCode, data, global.GqaConfig.System.ErrorMessage, c)
}

func ResponseErrorMessageData(data interface{}, message string, c *gin.Context) {
	Result(global.GqaConfig.System.ErrorCode, data, message, c)
}

type ResponsePage struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type ResponsePageWithParentId struct {
	Records    interface{} `json:"records"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	ParentCode string      `json:"parentCode"`
}

func OrderByColumn(sortBy string, desc bool) interface{} {
	return clause.OrderByColumn{Column: clause.Column{Name: sortBy}, Desc: desc}
}
