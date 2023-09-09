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
			global.GqaSLogger.Error(utils.GqaI18n(c, "GetUserRoleFailed"))
			model.ResponseErrorMessage(utils.GqaI18n(c, "GetUserRoleFailed"), c)
			c.Abort()
			return
		}
		apiPath := c.Request.URL.RequestURI()
		apiMethod := c.Request.Method
		if len(roleList) == 0 {
			global.GqaSLogger.Error(utils.GqaI18n(c, "NoRoleForYou"))
			model.ResponseErrorMessage(utils.GqaI18n(c, "NoRoleForYou"), c)
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
		global.GqaSLogger.Error(fmt.Sprintf("【%s】 tries to call 【%s】(%s), but does not have permission!", username, apiPath, apiMethod))
		model.ResponseErrorMessage(utils.GqaI18n(c, "NoPermission"), c)
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
