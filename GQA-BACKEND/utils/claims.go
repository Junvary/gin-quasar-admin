package utils

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
)

func GetUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GqaLog.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*system.GqaJwtClaims)
		return waitUse.Username
	}
}
