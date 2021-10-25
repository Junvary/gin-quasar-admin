package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Result(code int, data interface{}, message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

func SuccessMessage(message string, c *gin.Context) {
	Result(GqaConfig.System.SuccessCode, map[string]interface{}{}, message, c)
}

func SuccessData(data interface{}, c *gin.Context) {
	Result(GqaConfig.System.SuccessCode, data, GqaConfig.System.SuccessMessage, c)
}

func SuccessMessageData(data interface{}, message string, c *gin.Context) {
	Result(GqaConfig.System.SuccessCode, data, message, c)
}

func ErrorMessage(message string, c *gin.Context) {
	Result(GqaConfig.System.ErrorCode, map[string]interface{}{}, message, c)
}

func ErrorData(data interface{}, c *gin.Context) {
	Result(GqaConfig.System.ErrorCode, data, GqaConfig.System.ErrorMessage, c)
}

func ErrorMessageData(data interface{}, message string, c *gin.Context) {
	Result(GqaConfig.System.ErrorCode, data, message, c)
}
