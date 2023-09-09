package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var PluginAchievementSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-achievement").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_api 表中achievement插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> sys_api 表中achievement插件数据已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> achievement插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> achievement插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件achievement：获取Category-list",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-category-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：编辑Category信息",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/edit-category"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：新增Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/add-category"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：删除Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/delete-category-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：根据ID查找Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/query-category-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 6, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件achievement：获取Obtain-list",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-obtain-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 7, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：查找是否获取了成就",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/obtain-find"},
}
