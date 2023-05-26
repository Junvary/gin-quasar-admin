package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysButton = new(sysButton)

type sysButton struct{}

func (s *sysButton) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysButton{}).Count(&count)
		if count != 0 {
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_button"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_button"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysButtonData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_button"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_button"))
		return nil
	})
}

var sysButtonData = []model.SysButton{
	{MenuName: "dept", ButtonName: "新增父级部门", ButtonCode: "dept:addParent"},
	{MenuName: "dept", ButtonName: "编辑", ButtonCode: "dept:edit"},
	{MenuName: "dept", ButtonName: "部门用户", ButtonCode: "dept:deptUser"},
	{MenuName: "dept", ButtonName: "新增子级部门", ButtonCode: "dept:addChildren"},
	{MenuName: "dept", ButtonName: "删除", ButtonCode: "dept:delete"},

	{MenuName: "menu", ButtonName: "新增父级菜单", ButtonCode: "menu:addParent"},
	{MenuName: "menu", ButtonName: "编辑", ButtonCode: "menu:edit"},
	{MenuName: "menu", ButtonName: "新增子级部门", ButtonCode: "menu:addChildren"},
	{MenuName: "menu", ButtonName: "按钮注册", ButtonCode: "menu:buttonRegister"},
	{MenuName: "menu", ButtonName: "删除", ButtonCode: "menu:delete"},

	{MenuName: "dict", ButtonName: "新增父级字典", ButtonCode: "dict:addParent"},
	{MenuName: "dict", ButtonName: "编辑", ButtonCode: "dict:edit"},
	{MenuName: "dict", ButtonName: "新增子级字典", ButtonCode: "dict:addChildren"},
	{MenuName: "dict", ButtonName: "删除", ButtonCode: "dict:delete"},

	{MenuName: "api", ButtonName: "新增Api", ButtonCode: "api:add"},
	{MenuName: "api", ButtonName: "编辑", ButtonCode: "api:edit"},
	{MenuName: "api", ButtonName: "删除", ButtonCode: "api:delete"},

	{MenuName: "role", ButtonName: "新增角色", ButtonCode: "role:add"},
	{MenuName: "role", ButtonName: "编辑", ButtonCode: "role:edit"},
	{MenuName: "role", ButtonName: "角色用户", ButtonCode: "role:user"},
	{MenuName: "role", ButtonName: "角色权限", ButtonCode: "role:permission"},
	{MenuName: "role", ButtonName: "删除", ButtonCode: "role:delete"},

	{MenuName: "user", ButtonName: "新增用户", ButtonCode: "user:add"},
	{MenuName: "user", ButtonName: "重置密码", ButtonCode: "user:password"},
	{MenuName: "user", ButtonName: "编辑", ButtonCode: "user:edit"},
	{MenuName: "user", ButtonName: "删除", ButtonCode: "user:delete"},

	{MenuName: "config-frontend", ButtonName: "新增配置", ButtonCode: "config-frontend:add"},
	{MenuName: "config-frontend", ButtonName: "保存", ButtonCode: "config-frontend:save"},
	{MenuName: "config-frontend", ButtonName: "重置", ButtonCode: "config-frontend:reset"},
	{MenuName: "config-frontend", ButtonName: "删除", ButtonCode: "config-frontend:delete"},

	{MenuName: "config-backend", ButtonName: "新增配置", ButtonCode: "config-backend:add"},
	{MenuName: "config-backend", ButtonName: "保存", ButtonCode: "config-backend:save"},
	{MenuName: "config-backend", ButtonName: "重置", ButtonCode: "config-backend:reset"},
	{MenuName: "config-backend", ButtonName: "删除", ButtonCode: "config-backend:delete"},

	{MenuName: "log-login", ButtonName: "详情", ButtonCode: "log-login:detail"},
	{MenuName: "log-login", ButtonName: "删除", ButtonCode: "log-login:delete"},

	{MenuName: "log-operation", ButtonName: "详情", ButtonCode: "log-operation:detail"},
	{MenuName: "log-operation", ButtonName: "删除", ButtonCode: "log-operation:delete"},

	{MenuName: "notice", ButtonName: "新增消息", ButtonCode: "notice:add"},
	{MenuName: "notice", ButtonName: "发送", ButtonCode: "notice:send"},
	{MenuName: "notice", ButtonName: "编辑", ButtonCode: "notice:edit"},
	{MenuName: "notice", ButtonName: "删除", ButtonCode: "notice:delete"},

	{MenuName: "genPlugin", ButtonName: "新增插件", ButtonCode: "genPlugin:add"},
	{MenuName: "genPlugin", ButtonName: "生成插件", ButtonCode: "genPlugin:gen"},
	{MenuName: "genPlugin", ButtonName: "下载插件", ButtonCode: "genPlugin:download"},
	{MenuName: "genPlugin", ButtonName: "下载插件", ButtonCode: "genPlugin:delete"},

	{MenuName: "user-online", ButtonName: "踢出", ButtonCode: "user-online:kick"},
}
