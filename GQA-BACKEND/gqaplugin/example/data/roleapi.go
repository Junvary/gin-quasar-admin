package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var PluginExampleSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-example").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_api 表中example插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-Plugin] --> sys_role_api 表中example插件数据已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> example插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-Plugin] --> example插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/get-test-data-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/edit-test-data"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/add-test-data"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/delete-test-data-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/query-test-data-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/download-template-test-data"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/export-test-data"},
	{RoleCode: "super-admin", ApiGroup: "plugin-example", ApiMethod: "POST", ApiPath: "/plugin-example/import-test-data"},
}
