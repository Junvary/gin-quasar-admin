package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var PluginAchievementSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-achievement").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_role_api 表中Achievement插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> sys_role_api 表中Achievement插件数据已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> Achievement插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> Achievement插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{

	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-category-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/edit-category"},
	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/add-category"},
	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/delete-category-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/query-category-by-id"},

	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/get-obtain-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-achievement", ApiMethod: "POST", ApiPath: "/plugin-achievement/obtain-find"},
}
