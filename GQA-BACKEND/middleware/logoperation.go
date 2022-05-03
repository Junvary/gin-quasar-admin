package middleware

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

func LogOperationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var logOperation model.SysLogOperation
		logOperation.OperationUsername = utils.GetUsername(c)
		logOperation.OperationIp = c.ClientIP()
		logOperation.OperationApi = c.Request.RequestURI
		logOperation.OperationMethod = c.Request.Method
		logOperation.OperationStatus = c.Writer.Status()

		_ = global.GqaDb.Create(&logOperation).Error

		c.Next()
	}
}
