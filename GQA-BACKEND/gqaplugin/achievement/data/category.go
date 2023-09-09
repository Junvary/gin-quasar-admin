package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var PluginAchievementCategory = new(category)

type category struct{}

func (s *category) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&model.PluginAchievementCategory{}).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> plugin_achievement_category 表中已存在数据，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> plugin_achievement_category 表中已存在数据，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&categoryData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> Achievement插件初始数据进入 plugin_achievement_category 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> Achievement插件初始数据进入 plugin_achievement_category 表成功！")
		return nil
	})
}

var categoryData = []model.PluginAchievementCategory{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "发现某个可以旋转的图标",
	}}, Category: "QiYu", Code: "QiYu-Roll-Icon", Name: "摇滚吧，少年！"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "发现404页面",
	}}, Category: "QiYu", Code: "QiYu-Find-404", Name: "未知领域！"},
}
