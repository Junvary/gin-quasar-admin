package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginXkSysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Status: "on", Sort: 801, Remark: "这是系统开发科插件", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "GqaPluginXk", Title: "系统开发科", Icon: "house",
		Path: "", Component: "",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是最新要闻", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-xk-news", ParentCode: "GqaPluginXk", Title: "最新要闻", Icon: "newspaper",
		Path: "/plugin-xk/xk/news", Component: "/Plugin/Xk/News/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 2, Remark: "这是项目管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-xk-project", ParentCode: "GqaPluginXk", Title: "项目管理", Icon: "beenhere",
		Path: "/plugin-xk/xk/project", Component: "/Plugin/Xk/Project/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "这是荣誉认证", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-xk-honour", ParentCode: "GqaPluginXk", Title: "荣誉认证", Icon: "stars",
		Path: "/plugin-xk/xk/honour", Component: "/Plugin/Xk/Honour/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 4, Remark: "这是文档管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-xk-document", ParentCode: "GqaPluginXk", Title: "文档管理", Icon: "edit_note",
		Path: "/plugin-xk/xk/document", Component: "/Plugin/Xk/Document/index",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 5, Remark: "这是资源管理", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-xk-download", ParentCode: "GqaPluginXk", Title: "资源管理", Icon: "download",
		Path: "/plugin-xk/xk/download", Component: "/Plugin/Xk/Download/index",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Where("parent_code = ?", "GqaPluginXk").Or("name = ?", "GqaPluginXk").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中xk插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_menu 表中xk插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_menu 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> xk插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}
