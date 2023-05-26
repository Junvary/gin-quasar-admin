package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var Plugin{{.PluginCode}}SysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPlugin{{.PluginCode}}").Or("name = ?", "GqaPlugin{{.PluginCode}}").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_menu 表中{{.PluginCode}}插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-plugins] --> sys_menu 表中{{.PluginCode}}插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是{{ .PluginName }}插件",
	}}, Name: "GqaPlugin{{.PluginCode}}", Title: "{{ .PluginName }}", Icon: "home", Path: "", Component: ""},
	{{ range .PluginModel }}
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是{{ .ModelName }}",
	}}, Name: "plugin-{{$.PluginCode}}-{{ .ModelName }}", Title: "{{ .ModelName }}", Icon: "home", Path: "/plugin-{{$.PluginCode}}/{{$.PluginCode}}/{{ .ModelName }}", Component: "plugins/{{$.PluginCode}}/{{ .ModelName }}/index", ParentCode: "GqaPlugin{{$.PluginCode}}"},
    {{ end }}
}
