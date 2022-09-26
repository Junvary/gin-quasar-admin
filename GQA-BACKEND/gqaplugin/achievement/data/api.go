package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAchievementSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-achievement").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_api 表中achievement插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-plugins] --> sys_api 表中achievement插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> achievement插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-plugins] --> achievement插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70001, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件achievement：获取Category-list",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-category-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70002, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：编辑Category信息",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/edit-category"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70003, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：新增Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/add-category"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70004, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：删除Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/delete-category-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70005, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：根据ID查找Category",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/query-category-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70001, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件achievement：获取Obtain-list",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-obtain-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 70002, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件Achievement：查找是否获取了成就",
	}}, ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/obtain-find"},
}
