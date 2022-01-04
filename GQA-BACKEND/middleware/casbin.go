package middleware

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		err, subList := utils.GetUserRole(c)
		if err != nil {
			global.GqaLog.Error("Casbin中间件获取角色失败！")
			global.ErrorMessage("Casbin中间件获取角色失败！", c)
			c.Abort()
			return
		}
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		if len(subList) == 0 {
			global.GqaLog.Error("没有分配角色，你没有权限！")
			global.ErrorMessage("没有分配角色，你没有权限！", c)
			c.Abort()
			return
		}
		for _, v := range subList {
			success, _ := global.GqaCasbin.Enforce(v.RoleCode, obj, act)
			if success {
				c.Next()
				return
			}
		}
		username := utils.GetUsername(c)
		global.GqaLog.Error(fmt.Sprintf("【%s】 尝试通过【%s】 调用 【%s】 ，但没有权限！", username, act, obj))
		global.ErrorMessage("对不起，你没有权限！", c)
		c.Abort()
		return

	}
}
