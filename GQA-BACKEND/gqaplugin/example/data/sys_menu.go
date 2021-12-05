package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Status: "on", Sort: 501, Remark: "这是插件菜单示例", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "GqaPluginExample", Title: "插件示例", Icon: "install_desktop",
		Path: "", Component: "",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是示例插件", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "plugin-example", ParentCode: "GqaPluginExample", Title: "example插件", Icon: "install_desktop",
		Path: "/plugin-example/example", Component: "/Plugin/Example/index",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Where("parent_code = ?", "GqaPluginExample").Or("name = ?", "GqaPluginExample").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件初始数据进入 sys_menu 表成功！")
		global.GqaLog.Error("[GQA-Plugin] --> example插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}
