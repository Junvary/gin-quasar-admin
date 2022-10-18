package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginExampleSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-example").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中example插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_api 表中example插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> example插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 101, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：获取test-data-list",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/get-test-data-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 102, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：编辑test-data信息",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/edit-test-data"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 103, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：新增test-data",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/add-test-data"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 104, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：删除test-data",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/delete-test-data-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 105, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：根据ID查找test-data",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/query-test-data-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 106, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：下载test-data模板",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/download-template-test-data"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 107, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：导出test-data",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/export-test-data"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 108, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件example：导入test-data",
	}}, ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/import-test-data"},
}
