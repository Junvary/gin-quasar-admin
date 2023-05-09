package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
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
			fmt.Println(utils.GqaI18nWithData("SkipInsertWithData", "sys_menu"), count)
			global.GqaLogger.Warn(utils.GqaI18nWithData("SkipInsertWithData", "sys_menu"), zap.Any("count", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println(utils.GqaI18nWithData("TableInitSuccess", "sys_menu"))
		global.GqaLogger.Info(utils.GqaI18nWithData("TableInitSuccess", "sys_menu"))
		return nil
	})
}

var sysMenuData = []model.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: GqaSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "系统管理",
	}}, Name: "system", Title: "SystemManage", Icon: "eva-settings-2-outline", Path: "", Component: "", Redirect: "/dashboard"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "仪表盘",
	}}, Name: "dashboard", Title: "Dashboard", Icon: "eva-home-outline", Path: "/dashboard", Component: "pages/Dashboard/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "部门管理",
	}}, Name: "dept", Title: "DeptManage", Icon: "eva-grid-outline", Path: "/dept", Component: "pages/System/Dept/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "菜单管理",
	}}, Name: "menu", Title: "MenuManage", Icon: "eva-menu", Path: "/menu", Component: "pages/System/Menu/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "字典管理",
	}}, Name: "dict", Title: "DictManage", Icon: "eva-book-open-outline", Path: "/dict", Component: "pages/System/Dict/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Api管理",
	}}, Name: "api", Title: "ApiManage", Icon: "eva-link", Path: "/api", Component: "pages/System/Api/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "角色管理",
	}}, Name: "role", Title: "RoleManage", Icon: "eva-people-outline", Path: "/role", Component: "pages/System/Role/index", ParentCode: "system"},
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
	}}, Name: "log", Title: "LogManage", Icon: "eva-printer-outline", Path: "/log", Component: "", Redirect: "/log/log-login", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "登录日志",
	}}, Name: "log-login", Title: "LogLogin", Icon: "eva-printer-outline", Path: "/log/log-login", Component: "pages/System/Log/Login/index", ParentCode: "log"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "操作日志",
	}}, Name: "log-operation", Title: "LogOperation", Icon: "eva-printer-outline", Path: "/log/log-operation", Component: "pages/System/Log/Operation/index", ParentCode: "log"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 11, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "消息管理",
	}}, Name: "notice", Title: "NoticeManage", Icon: "eva-message-circle-outline", Path: "/notice", Component: "pages/System/Notice/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 12, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件生成",
	}}, Name: "genPlugin", Title: "GenPlugin", Icon: "eva-code-download", Path: "/gen-plugin", Component: "pages/System/GenPlugin/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 13, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "在线用户",
	}}, Name: "user-online", Title: "UserOnline", Icon: "eva-monitor-outline", Path: "/user-online", Component: "pages/System/UserOnline/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 14, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "定时任务",
	}}, Name: "cron", Title: "Cron", Icon: "eva-clock-outline", Path: "/cron", Component: "pages/System/Cron/index", ParentCode: "system"},
	{GqaModelWithCreatedByAndUpdatedBy: model.GqaModelWithCreatedByAndUpdatedBy{GqaModel: global.GqaModel{
		Sort: 100, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "Github",
	}}, Name: "github", Title: "Github", Icon: "eva-github", IsLink: "yesNo_yes", Path: "https://github.com/Junvary/gin-quasar-admin", Component: "", ParentCode: "system"},
}
