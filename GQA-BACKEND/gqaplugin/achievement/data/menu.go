package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var PluginAchievementSysMenu = new(sysMenu)

type sysMenu struct{}

func (s *sysMenu) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysMenu{}).Where("parent_code = ?", "GqaPluginAchievement").Or("name = ?", "GqaPluginAchievement").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_menu 表中Achievement插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> sys_menu 表中Achievement插件菜单已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&sysMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> Achievement插件初始数据进入 sys_menu 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> Achievement插件初始数据进入 sys_menu 表成功！")
		return nil
	})
}

var sysMenuData = []gqaModel.SysMenu{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是成就系统插件",
	}}, Name: "GqaPluginAchievement", Title: "成就插件", Icon: "fas fa-award", Path: "", Component: "", Redirect: "/plugin-achievement/achievement/category"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是Category",
	}}, Name: "plugin-achievement-category", Title: "成就分类", Icon: "fas fa-award", Path: "/plugin-achievement/achievement/category", Component: "plugins/Achievement/Category/index", ParentCode: "GqaPluginAchievement"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "这是Obtain",
	}}, Name: "plugin-achievement-obtain", Title: "已获得成就", Icon: "fas fa-award", Path: "/plugin-achievement/achievement/obtain", Component: "plugins/Achievement/Obtain/index", ParentCode: "GqaPluginAchievement"},
}
