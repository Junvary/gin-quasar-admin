package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
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

func (s *sysRoleMenu)Init()error  {
	return global.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_role_id IN ('1')").Find(&[]system.SysRoleMenu{}).RowsAffected == 8 {
			fmt.Println("\n[Mysql] --> sys_role_menu 表的初始数据已存在！")
			global.GqaLog.Error("sys_role_menu 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&sysRoleMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_role_menu 表初始数据成功!")
		global.GqaLog.Error("sys_role_menu 表初始数据成功!")
		return nil
	})
}
