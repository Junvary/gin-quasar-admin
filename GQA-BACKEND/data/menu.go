package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.SysMenu{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLogger.Warn("[Gin-Quasar-Admin] --> sys_menu 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表初始数据成功！")
		global.GqaLogger.Info("[Gin-Quasar-Admin] --> sys_menu 表初始数据成功！")
		return nil
	})
}

var sysMenuData = []model.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "系统管理",
	}}, Name: "system", Title: "SystemManage", Icon: "eva-settings-2-outline", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "仪表盘",
	}}, Name: "dashboard", Title: "Dashboard", Icon: "mdi-view-dashboard-outline", Path: "/dashboard", Component: "pages/Dashboard/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "部门管理",
	}}, Name: "dept", Title: "DeptManage", Icon: "work_outline", Path: "/dept", Component: "pages/System/Dept/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "菜单管理",
	}}, Name: "menu", Title: "MenuManage", Icon: "eva-menu", Path: "/menu", Component: "pages/System/Menu/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "字典管理",
	}}, Name: "dict", Title: "DictManage", Icon: "ion-ios-list", Path: "/dict", Component: "pages/System/Dict/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Api管理",
	}}, Name: "api", Title: "ApiManage", Icon: "ion-ios-infinite", Path: "/api", Component: "pages/System/Api/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "角色管理",
	}}, Name: "role", Title: "RoleManage", Icon: "people_outline", Path: "/role", Component: "pages/System/Role/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "用户管理",
	}}, Name: "user", Title: "UserManage", Icon: "eva-person-outline", Path: "/user", Component: "pages/System/User/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 8, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "前台配置管理",
	}}, Name: "config-frontend", Title: "FrontendManage", Icon: "eva-settings-outline", Path: "/config-frontend", Component: "pages/System/ConfigFrontend/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 9, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "后台配置管理",
	}}, Name: "config-backend", Title: "BackendManage", Icon: "eva-settings-outline", Path: "/config-backend", Component: "pages/System/ConfigBackend/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 10, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "系统日志",
	}}, Name: "log", Title: "LogManage", Icon: "ion-md-book", Path: "/log", Component: "", Redirect: "/log/log-login", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录日志",
	}}, Name: "log-login", Title: "LogLogin", Icon: "ion-md-book", Path: "/log/log-login", Component: "pages/System/Log/Login/index", ParentCode: "log"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "操作日志",
	}}, Name: "log-operation", Title: "LogOperation", Icon: "ion-md-book", Path: "/log/log-operation", Component: "pages/System/Log/Operation/index", ParentCode: "log"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 11, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息管理",
	}}, Name: "notice", Title: "NoticeManage", Icon: "eva-message-circle-outline", Path: "/notice", Component: "pages/System/Notice/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 12, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件生成",
	}}, Name: "genPlugin", Title: "GenPlugin", Icon: "ion-md-code-working", Path: "/gen-plugin", Component: "pages/System/GenPlugin/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 13, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "在线用户",
	}}, Name: "user-online", Title: "UserOnline", Icon: "eva-people-outline", Path: "/user-online", Component: "pages/System/UserOnline/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 14, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Github",
	}}, Name: "github", Title: "Github", Icon: "fab fa-github", IsLink: "yes", Path: "https://github.com/Junvary/gin-quasar-admin", Component: "", ParentCode: "system"},
}
