package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginExample").Or("name = ?", "GqaPluginExample").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> example插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1101, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是示例插件",
	}}, IsPlugin: "yes", Name: "GqaPluginExample", Title: "示例插件", Icon: "install_desktop", Path: "", Component: ""},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是图标合集",
	}}, IsPlugin: "yes", Name: "plugin-example-icons", Title: "图标合集", Icon: "insert_emoticon", Path: "/plugin-example/example/icons", Component: "plugins/Example/Icons/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是编辑器",
	}}, IsPlugin: "yes", Name: "plugin-example-editor", Title: "编辑器", Icon: "edit", Path: "/plugin-example/example/editor", Component: "plugins/Example/Editor/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是树形表格",
	}}, IsPlugin: "yes", Name: "plugin-example-tree-table", Title: "树形表格", Icon: "list_alt", Path: "/plugin-example/example/tree-table", Component: "plugins/Example/TreeTable/index", ParentCode: "GqaPluginExample"},
}
