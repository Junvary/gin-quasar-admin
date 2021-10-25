package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
	"time"
)

var SysRole = new(sysRole)

type sysRole struct{}

var roles = []system.SysRole{
	{
		GqaModel: global.GqaModel{
			Id:       1,
			Status:   "on",
			Sort:     1,
			Desc:     "超级管理员角色",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		RoleCode: "super-admin",
		RoleName: "超级管理组",
	},
}

func (a *sysRole) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysRole{}).RowsAffected == 1 {
			fmt.Println("\n[Mysql] --> sys_role 表的初始数据已存在！")
			global.GqaLog.Error("sys_role 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&roles).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_role 表初始数据成功！")
		global.GqaLog.Error("sys_role 表初始数据成功！")
		return nil
	})
}
