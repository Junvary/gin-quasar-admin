package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysRoleButton = new(sysRoleButton)

type sysRoleButton struct{}

func (s *sysRoleButton) LoadData() error {
	return global.GqaDb.Table("sys_role_button").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysRoleButton{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_role_button 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_role_button 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleButtonData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_role_button 表初始数据成功!")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_role_button 表初始数据成功!")
		return nil
	})
}

var sysRoleButtonData = []model.SysRoleButton{
	{"super-admin", "dept:addParent"},
	{"super-admin", "dept:edit"},
	{"super-admin", "dept:deptUser"},
	{"super-admin", "dept:addChildren"},
	{"super-admin", "dept:delete"},

	{"super-admin", "menu:addParent"},
	{"super-admin", "menu:edit"},
	{"super-admin", "menu:addChildren"},
	{"super-admin", "menu:buttonRegister"},
	{"super-admin", "menu:delete"},

	{"super-admin", "dict:addParent"},
	{"super-admin", "dict:edit"},
	{"super-admin", "dict:addChildren"},
	{"super-admin", "dict:delete"},

	{"super-admin", "api:add"},
	{"super-admin", "api:edit"},
	{"super-admin", "api:delete"},

	{"super-admin", "role:add"},
	{"super-admin", "role:edit"},
	{"super-admin", "role:user"},
	{"super-admin", "role:permission"},
	{"super-admin", "role:delete"},

	{"super-admin", "user:add"},
	{"super-admin", "user:password"},
	{"super-admin", "user:edit"},
	{"super-admin", "user:delete"},

	{"super-admin", "config-frontend:add"},
	{"super-admin", "config-frontend:save"},
	{"super-admin", "config-frontend:reset"},
	{"super-admin", "config-frontend:delete"},

	{"super-admin", "config-backend:add"},
	{"super-admin", "config-backend:save"},
	{"super-admin", "config-backend:reset"},
	{"super-admin", "config-backend:delete"},

	{"super-admin", "log-login:detail"},
	{"super-admin", "log-login:delete"},

	//{"super-admin", "log-operation:detail"},
	//{"super-admin", "log-operation:delete"},

	{"super-admin", "notice:add"},
	{"super-admin", "notice:send"},
	{"super-admin", "notice:delete"},

	{"super-admin", "genPlugin:gen"},

	{"super-admin", "user-online:kick"},
}
