package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysUserRole = new(sysUserRole)

type sysUserRole struct{}

func (s *sysUserRole) LoadData() error {
	return global.GqaDb.Table("sys_user_role").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysUserRole{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_user_role"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_user_role"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysUserRoleData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_user_role"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_user_role"))
		return nil
	})
}

var sysUserRoleData = []model.SysUserRole{
	{"super-admin", "admin"},
}
