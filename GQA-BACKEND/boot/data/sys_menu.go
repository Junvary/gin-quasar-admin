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
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 1, Stable: "yes", Status: "on", Sort: 1, Remark: "这是仪表盘", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 0, Name: "dashboard", Path: "/dashboard", Component: "/Dashboard/index",
		Title: "仪表盘", Icon: "home", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 2, Stable: "yes", Status: "on", Sort: 2, Remark: "这是系统管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 0, Name: "system", Path: "", Component: "",
		Title: "系统管理", Icon: "settings", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	// 父级设置ID
	{GqaModel: global.GqaModel{Id: 3, Stable: "yes", Status: "on", Sort: 3, Remark: "这是系统示例", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 0, Name: "example", Path: "", Component: "",
		Title: "系统示例", Icon: "star", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},

	// 子级内容：
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是部门管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "dept", Path: "/system/dept", Component: "/System/Dept/index",
		Title: "部门管理", Icon: "work", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 2, Remark: "这是用户管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "user", Path: "/system/user", Component: "/System/User/index",
		Title: "用户管理", Icon: "person", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 3, Remark: "这是角色管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "role", Path: "/system/role", Component: "/System/Role/index",
		Title: "角色管理", Icon: "group", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 4, Remark: "这是菜单管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "menu", Path: "/system/menu", Component: "/System/Menu/index",
		Title: "菜单管理", Icon: "menu", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 5, Remark: "这是字典管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "dict", Path: "/system/dict", Component: "/System/Dict/index",
		Title: "字典管理", Icon: "zoom_in", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 6, Remark: "这是配置管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "config", Path: "/system/config", Component: "/System/Config/index",
		Title: "配置管理", Icon: "settings", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 7, Remark: "这是日志管理", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 2, Name: "log", Path: "/system/log", Component: "/System/Log/index",
		Title: "日志管理", Icon: "toc", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},

	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1, Remark: "这是图标合集", CreateAt: time.Now(), CreateBy: "admin", UpdateAt: time.Now()},
		ParentId: 3, Name: "icon", Path: "/example/icon", Component: "/Example/Icon/index",
		Title: "图标合集", Icon: "mood", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
}

func (s *sysMenu) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("sys_menu 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_menu 表初始数据成功！")
		global.GqaLog.Error("sys_menu 表初始数据成功！")
		return nil
	})
}
