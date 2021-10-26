package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct {
}

var sysRoleMenus = []system.SysRoleMenu{
	{"super-admin", 1},
	{"super-admin", 2},
	{"super-admin", 3},
	{"super-admin", 4},
	{"super-admin", 5},
	{"super-admin", 6},
	{"super-admin", 7},
	{"super-admin", 8},
}

func (s *sysRoleMenu) Init() error {
	return global.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysRoleMenu{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_role_menu 表的初始数据已存在！数据量：", count)
			global.GqaLog.Error("sys_role_menu 表的初始数据已存在！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_role_menu 表初始数据成功!")
		global.GqaLog.Error("sys_role_menu 表初始数据成功!")
		return nil
	})
}
