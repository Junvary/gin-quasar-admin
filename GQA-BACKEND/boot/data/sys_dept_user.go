package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysDeptUser = new(sysDeptUser)

type sysDeptUser struct{}

var sysDeptUserData = []system.SysDeptUser{
	{"gin-quasar-admin", "admin"},
}

func (s *sysDeptUser) LoadData() error {
	return global.GqaDb.Table("sys_dept_user").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysDeptUser{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_dept_user 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_dept_user 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDeptUserData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_dept_user 表初始数据成功!")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_dept_user 表初始数据成功!")
		return nil
	})
}
