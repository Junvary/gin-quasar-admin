package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var SysUserRole = new(sysUserRole)

type sysUserRole struct{}

func (s *sysUserRole) LoadData(c *gin.Context) error {
	return global.GqaDb.Table("sys_user_role").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysUserRole{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_user_role"), count)
			global.GqaSLogger.Warn(utils.GqaI18nWithData(c, "SkipInsertWithData", "sys_user_role"), "has_count", count)
			return nil
		}
		if err := tx.Create(&sysUserRoleData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_user_role"))
		global.GqaSLogger.Info(utils.GqaI18nWithData(c, "TableInitSuccess", "sys_user_role"))
		return nil
	})
}

var sysUserRoleData = []model.SysUserRole{
	{"super-admin", "admin"},
}
