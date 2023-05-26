package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Plugin{{.PluginCode}}SysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-{{.PluginCode}}").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_role_api 表中{{.PluginCode}}插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-plugins] --> sys_role_api 表中{{.PluginCode}}插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-plugins] --> {{.PluginCode}}插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
    {{ range .PluginModel }}
	{RoleCode: "super-admin", ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/get-{{.ModelName}}-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/edit-{{.ModelName}}"},
	{RoleCode: "super-admin", ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/add-{{.ModelName}}"},
	{RoleCode: "super-admin", ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/delete-{{.ModelName}}-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-{{$.PluginCode}}", ApiMethod: "POST", ApiPath: "/plugin-{{$.PluginCode}}/query-{{.ModelName}}-by-id"},
	{{ end }} 
}
