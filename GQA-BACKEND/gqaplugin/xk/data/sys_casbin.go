package data

import (
	"fmt"
	"gin-quasar-admin/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginXkSysCasbin = new(sysCasbin)

type sysCasbin struct{}

var sysCasbinData = []gormadapter.CasbinRule{
	// plugin-xk组
	//news
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/news-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/news-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/news-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/news-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/news-id", V2: "POST"},
	//project
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/project-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/project-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/project-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/project-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/project-id", V2: "POST"},
	//honour
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/honour-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/honour-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/honour-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/honour-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/honour-id", V2: "POST"},
	//document
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/document-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/document-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/document-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/document-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/document-id", V2: "POST"},
	//download
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/download-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/download-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/download-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/download-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-xk/download-id", V2: "POST"},
}

func (s *sysCasbin) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&[]gormadapter.CasbinRule{}).Where("v1 like ?", "/plugin-xk%").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> casbin_rule 表中xk插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> casbin_rule 表中xk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysCasbinData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 casbin_rule 表成功！")
		global.GqaLog.Error("[GQA-Plugin] --> xk插件初始数据进入 casbin_rule 表成功！")
		return nil
	})
}
