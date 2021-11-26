package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysMenu = new(sysMenu)

type sysMenu struct{}

var sysMenuData = []system.SysMenu{
	{GqaModel: global.GqaModel{Status: "on", Sort: 3, Remark: "这是插件菜单示例", CreatedAt: time.Now(), CreatedBy: "admin"},
		Name: "GqaPlugin", Path: "", Component: "",
		Title: "插件示例", Icon: "install_desktop", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
	{GqaModel: global.GqaModel{Status: "on", Sort: 1, Remark: "这是示例插件", CreatedAt: time.Now(), CreatedBy: "admin"},
		ParentCode: "GqaPlugin", Name: "plugin-example", Path: "/plugin-example/example", Component: "/Plugin/Example/index",
		Title: "example插件", Icon: "install_desktop", Hidden: "no", KeepAlive: "no", IsLink: "no",
	},
}

func (s *sysMenu) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Where("parent_code = ?", "GqaPlugin").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件数据进入 sys_menu 表初始数据成功！")
		global.GqaLog.Error("[GQA-Plugin] --> example插件数据进入 sys_menu 表初始数据成功！")
		return nil
	})
}
