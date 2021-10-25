package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gorm.io/gorm"
	"time"
)

var SysMenu = new(sysMenu)

type sysMenu struct{}

var menus = []system.SysMenu{
	{
		GqaModel: global.GqaModel{
			Id:       1,
			Status:   "on",
			Sort:     1,
			Desc:     "这是仪表盘",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  0,
		Path:      "/dashboard",
		Component: "/Dashboard/index",
		Hidden: false,
		KeepAlive: false,
		Title: "仪表盘",
		Icon:  "home",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       2,
			Status:   "on",
			Sort:     2,
			Desc:     "这是系统管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  0,
		Path:      "",
		Component: "",
		Hidden: false,
		KeepAlive: false,
		Title: "系统管理",
		Icon:  "settings",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       3,
			Status:   "on",
			Sort:     1,
			Desc:     "这是部门管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/dept",
		Component: "/System/Dept/index",
		Hidden: false,
		KeepAlive: false,
		Title: "部门管理",
		Icon:  "work",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       4,
			Status:   "on",
			Sort:     2,
			Desc:     "这是用户管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/user",
		Component: "/System/User/index",
		Hidden: false,
		KeepAlive: false,
		Title: "用户管理",
		Icon:  "person",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       5,
			Status:   "on",
			Sort:     3,
			Desc:     "这是角色管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/role",
		Component: "/System/Role/index",
		Hidden: false,
		KeepAlive: false,
		Title: "角色管理",
		Icon:  "group",
		IsLink: false,
	},

	{
		GqaModel: global.GqaModel{
			Id:       6,
			Status:   "on",
			Sort:     4,
			Desc:     "这是菜单管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/menu",
		Component: "/System/Menu/index",
		Hidden: false,
		KeepAlive: false,
		Title: "菜单管理",
		Icon:  "menu",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       7,
			Status:   "on",
			Sort:     5,
			Desc:     "这是字典管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/dict",
		Component: "/System/Dict/index",
		Hidden: false,
		KeepAlive: false,
		Title: "字典管理",
		Icon:  "zoom_in",
		IsLink: false,
	},
	{
		GqaModel: global.GqaModel{
			Id:       8,
			Status:   "on",
			Sort:     6,
			Desc:     "这是日志管理",
			CreateAt: time.Now(),
			CreateBy: "admin",
			UpdateAt: time.Now(),
		},
		ParentId:  2,
		Path:      "/system/log",
		Component: "/System/Log/index",
		Hidden: false,
		KeepAlive: false,
		Title: "日志管理",
		Icon:  "toc",
		IsLink: false,
	},
}

func (a *sysMenu) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysMenu{}).RowsAffected == 8 {
			fmt.Println("\n[Mysql] --> sys_menu 表的初始数据已存在！")
			global.GqaLog.Error("sys_menu 表的初始数据已存在！")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("\n[Mysql] --> sys_menu 表初始数据成功！")
		global.GqaLog.Error("sys_menu 表初始数据成功！")
		return nil
	})
}
