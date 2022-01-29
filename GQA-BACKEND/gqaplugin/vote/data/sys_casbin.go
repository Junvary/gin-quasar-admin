package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginVoteSysCasbin = new(sysCasbin)

type sysCasbin struct{}

var sysCasbinData = []gormadapter.CasbinRule{
	// plugin-vote组
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-random-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-random-get", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-random-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-random-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/vote-handle", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/can-vote-dy", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/can-vote-gl", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/vote-result-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/vote-result-chart", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/voter-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/candidate-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/candidate-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/plugin-vote/candidate-delete", V2: "DELETE"},
}

func (s *sysCasbin) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&[]gormadapter.CasbinRule{}).Where("v1 like ?", "/plugin-vote%").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> casbin_rule 表中vote插件数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> casbin_rule 表中vote插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysCasbinData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 casbin_rule 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> vote插件初始数据进入 casbin_rule 表成功！")
		return nil
	})
}
