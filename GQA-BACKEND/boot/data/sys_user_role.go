package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
)

var SysUserRole = new(sysUserRole)

type sysUserRole struct {
}

var sysUserRoles = []system.SysUserRole{
	{"super-admin", 1},
}

func (s *sysUserRole)Init()error  {
	return global.GqaDb.Table("sys_user_role").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_user_id IN ('1')").Find(&[]system.SysUserRole{}).RowsAffected == 1 {
			fmt.Println("\n[Mysql] --> sys_user_role 表的初始数据已存在！")
			global.GqaLog.Error("sys_user_role 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&sysUserRoles).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_user_role 表初始数据成功!")
		global.GqaLog.Error("sys_user_role 表初始数据成功!")
		return nil
	})
}
