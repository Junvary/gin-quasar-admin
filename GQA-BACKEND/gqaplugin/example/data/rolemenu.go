package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var PluginExampleSysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct{}

func (s *sysRoleMenu) LoadData(c *gin.Context) error {
	return gqaGlobal.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		var menuName []string
		for _, v := range sysRoleMenuData {
			menuName = append(menuName, v.SysMenuName)
		}
		tx.Model(&gqaModel.SysRoleMenu{}).Where("sys_menu_name in ?", menuName).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-plugins] --> sys_role_menu 表中example插件数已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaSLogger.Warn("[GQA-plugins] --> sys_role_menu 表中example插件数据已存在，跳过初始化数据！", "has_count", count)
			return nil
		}
		if err := tx.Save(&sysRoleMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-plugins] --> example插件初始数据进入 sys_role_menu 表成功！")
		gqaGlobal.GqaSLogger.Info("[GQA-plugins] --> example插件初始数据进入 sys_role_menu 表成功！")
		return nil
	})
}

var sysRoleMenuData = []gqaModel.SysRoleMenu{
	{"super-admin", "GqaPluginExample"},
	{"super-admin", "plugin-example-icons"},
	{"super-admin", "plugin-example-editor"},
	{"super-admin", "plugin-example-tree-table"},
	{"super-admin", "plugin-example-statistic"},
	{"super-admin", "plugin-example-cascader"},
	{"super-admin", "plugin-example-import-export"},
	{"super-admin", "plugin-example-uploader"},
}
