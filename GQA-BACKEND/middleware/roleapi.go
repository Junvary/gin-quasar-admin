package middleware

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

func RoleApiHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		err, roleList := GetUserRole(c)
		if err != nil {
			global.GqaLogger.Error("获取用户角色失败！")
			model.ResponseErrorMessage("获取用户角色失败！", c)
			c.Abort()
			return
		}
		apiPath := c.Request.URL.RequestURI()
		apiMethod := c.Request.Method
		if len(roleList) == 0 {
			global.GqaLogger.Error("没有分配角色，你没有权限！")
			model.ResponseErrorMessage("没有分配角色，你没有权限！", c)
			c.Abort()
			return
		}
		for _, v := range roleList {
			var roleApi model.SysRoleApi
			result := global.GqaDb.Where("role_code = ? and api_path = ? and api_method = ?", v.RoleCode, apiPath, apiMethod).First(&roleApi)
			if result.RowsAffected != 0 {
				c.Next()
				return
			}
		}
		username := utils.GetUsername(c)
		global.GqaLogger.Error(fmt.Sprintf("【%s】 尝试【%s】 调用 【%s】 ，但没有权限！", username, apiMethod, apiPath))
		model.ResponseErrorMessage("对不起，你没有权限！", c)
		c.Abort()
		return

	}
}

func GetUserRole(c *gin.Context) (err error, role []model.SysRole) {
	username := utils.GetUsername(c)
	var user model.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	var userRole []model.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&userRole)
	if err != nil {
		return err, nil
	}
	return nil, userRole
}
