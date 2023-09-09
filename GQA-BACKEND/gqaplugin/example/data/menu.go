package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginExample").Or("name = ?", "GqaPluginExample").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> sys_menu 表中example插件菜单已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> example插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> example插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是示例插件",
	}}, Name: "GqaPluginExample", Title: "示例插件", Icon: "install_desktop", Path: "", Component: "", Redirect: "/plugin-example/example/icons"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是图标合集",
	}}, Name: "plugin-example-icons", Title: "图标合集", Icon: "insert_emoticon", Path: "/plugin-example/example/icons", Component: "plugins/Example/Icons/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是编辑器",
	}}, Name: "plugin-example-editor", Title: "编辑器", Icon: "edit", Path: "/plugin-example/example/editor", Component: "plugins/Example/Editor/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是树形表格",
	}}, Name: "plugin-example-tree-table", Title: "树形表格", Icon: "list_alt", Path: "/plugin-example/example/tree-table", Component: "plugins/Example/TreeTable/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是统计数据",
	}}, Name: "plugin-example-statistic", Title: "统计数据", Icon: "filter_9_plus", Path: "/plugin-example/example/statistic", Component: "plugins/Example/Statistic/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是级联选择器",
	}}, Name: "plugin-example-cascader", Title: "级联选择器", Icon: "call_split", Path: "/plugin-example/example/cascader", Component: "plugins/Example/Cascader/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是数据导入导出",
	}}, Name: "plugin-example-import-export", Title: "导入导出", Icon: "import_export", Path: "/plugin-example/example/import-export", Component: "plugins/Example/ExportData/index", ParentCode: "GqaPluginExample"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是上传组件",
	}}, Name: "plugin-example-uploader", Title: "上传器", Icon: "eva-cloud-upload-outline", Path: "/plugin-example/example/uploader", Component: "plugins/Example/Uploader/index", ParentCode: "GqaPluginExample"},
}
