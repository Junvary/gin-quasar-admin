package middleware

import (
	"bytes"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"io"
)

func LogOperationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var logOperation model.SysLogOperation

		body, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		logOperation.OperationUsername = utils.GetUsername(c)
		logOperation.OperationIp = c.ClientIP()
		logOperation.OperationApi = c.Request.RequestURI
		logOperation.OperationMethod = c.Request.Method
		logOperation.OperationStatus = c.Writer.Status()
		logOperation.OperationBody = string(body)

		_ = global.GqaDb.Create(&logOperation).Error

		c.Next()
	}
}
