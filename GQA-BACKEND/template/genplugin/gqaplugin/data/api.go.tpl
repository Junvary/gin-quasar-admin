package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var Plugin{{.PluginCode}}SysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-{{.PluginCode}}").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_api 表中{{.PluginCode}}插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-plugins] --> sys_api 表中{{.PluginCode}}插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
    {{ range .PluginModel }}
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 1, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件{{$.PluginCode}}：获取{{ .ModelName }}-list",
	}}, ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/get-{{ .ModelName }}-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 2, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件{{$.PluginCode}}：编辑{{ .ModelName }}信息",
	}}, ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/edit-{{ .ModelName }}"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 3, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件{{$.PluginCode}}：新增{{ .ModelName }}",
	}}, ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/add-{{ .ModelName }}"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 4, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件{{$.PluginCode}}：删除{{ .ModelName }}",
	}}, ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/delete-{{ .ModelName }}-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: PluginSort + 5, Stable: "yesNo_yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件{{$.PluginCode}}：根据ID查找{{ .ModelName }}",
	}}, ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/query-{{ .ModelName }}-by-id"},
	{{ end }}
}
