package middleware

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		err, role := service.GroupServiceApp.ServiceSystem.ServiceUser.GetUserRole(c)
		if err!= nil{
			global.ErrorMessage("对不起，获取用户角色失败，" + err.Error(), c)
			c.Abort()
			return
		}
		e := service.GroupServiceApp.ServiceSystem.ServiceCasbin.Casbin()
		for r := range role{
			success, _ := e.Enforce(role[r].RoleCode, obj, act)
			if success{
				c.Next()
				break
			}
		}
		global.ErrorMessage("对不起，你没有权限！", c)
		c.Abort()
		return
	}
}
