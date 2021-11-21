package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysUserRole = new(sysUserRole)

type sysUserRole struct{}

var sysUserRoleData = []system.SysUserRole{
	{"super-admin", 1},
}

func (s *sysUserRole) LoadData() error {
	return global.GqaDb.Table("sys_user_role").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysUserRole{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_user_role 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_user_role 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysUserRoleData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_user_role 表初始数据成功!")
		global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_user_role 表初始数据成功!")
		return nil
	})
}
