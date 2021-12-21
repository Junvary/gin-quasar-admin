package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是仪表盘", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "dashboard", Title: "Dashboard", Icon: "home",
		Path: "/dashboard", Component: "/Dashboard/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是系统管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "system", Title: "SystemManage", Icon: "settings",
		Path: "", Component: "",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是部门管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "dept", ParentCode: "system", Title: "DeptManage", Icon: "work",
		Path: "/system/dept", Component: "/System/Dept/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是用户管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "user", ParentCode: "system", Title: "UserManage", Icon: "person",
		Path: "/system/user", Component: "/System/User/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "这是角色管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "role", ParentCode: "system", Title: "RoleManage", Icon: "group",
		Path: "/system/role", Component: "/System/Role/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "这是菜单管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "menu", ParentCode: "system", Title: "MenuManage", Icon: "menu",
		Path: "/system/menu", Component: "/System/Menu/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "这是字典管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "dict", ParentCode: "system", Title: "DictManage", Icon: "zoom_in",
		Path: "/system/dict", Component: "/System/Dict/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "这是前台配置管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "config-frontend", ParentCode: "system", Title: "FrontendManage", Icon: "settings",
		Path: "/system/config-frontend", Component: "/System/ConfigFrontend/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "这是后台配置管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "config-backend", ParentCode: "system", Title: "BackendManage", Icon: "settings",
		Path: "/system/config-backend", Component: "/System/ConfigBackend/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 8, Remark: "这是日志管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "log", ParentCode: "system", Title: "LogManage", Icon: "toc",
		Path: "/system/log", Component: "/System/Log/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是登录日志", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "log-login", ParentCode: "log", Title: "LogLogin", Icon: "toc",
		Path: "/system/log/log-login", Component: "/System/Log/Login/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是操作日志", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "log-operation", ParentCode: "log", Title: "LogOperation", Icon: "toc",
		Path: "/system/log/log-operation", Component: "/System/Log/Operation/index",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 9, Remark: "这是系统示例", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "example", ParentCode: "system", Title: "SystemExample", Icon: "star",
		Path: "", Component: "",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是图标合集", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "example-icon", ParentCode: "example", Title: "ExampleIcon", Icon: "mood",
		Path: "/system/example/icon", Component: "/System/Example/Icon/index",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_menu 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表初始数据成功！")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_menu 表初始数据成功！")
		return nil
	})
}
