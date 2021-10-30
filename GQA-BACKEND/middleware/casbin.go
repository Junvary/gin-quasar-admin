package middleware

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*system.GqaJwtClaims)
		sub := waitUse.Username
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		success, _ := global.GqaCasbin.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			global.GqaLog.Error(fmt.Sprintf("【%s】 尝试通过【%s】 调用 【%s】 ，但没有权限！", sub, act, obj))
			global.ErrorMessage("对不起，你没有权限！", c)
			c.Abort()
			return
		}
	}
}
