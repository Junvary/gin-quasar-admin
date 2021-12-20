package middleware

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
)

func LogOperationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var logOperation system.SysLogOperation
		logOperation.OperationUsername = utils.GetUsername(c)
		logOperation.OperationIp = c.ClientIP()
		logOperation.OperationApi = c.Request.RequestURI
		logOperation.OperationMethod = c.Request.Method
		logOperation.OperationStatus = c.Writer.Status()

		_ = global.GqaDb.Create(&logOperation).Error

		c.Next()
	}
}
