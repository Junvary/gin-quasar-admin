package utils

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
)

func GetUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GqaLog.Error("解析token负载失败！")
		return ""
	} else {
		waitUse := claims.(*system.GqaJwtClaims)
		return waitUse.Username
	}
}
