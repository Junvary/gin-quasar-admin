package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginVoteSysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct{}

var sysRoleMenuData = []system.SysRoleMenu{
	// 为 super-admin 设置所有 sys_menu 的总数
	{"super-admin", "GqaPluginVote"},
	{"super-admin", "plugin-vote-vote-dy"},
	{"super-admin", "plugin-vote-vote-gl"},
	{"super-admin", "plugin-vote-voter-random"},
	{"super-admin", "plugin-vote-voter"},
	{"super-admin", "plugin-vote-candidate"},
	{"super-admin", "plugin-vote-score-result"},
}

func (s *sysRoleMenu) LoadData() error {
	return global.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Unscoped().Where("sys_menu_name in ?",
			[]string{"GqaPluginVote", "plugin-vote-news", "plugin-vote-project"}).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_menu 表中vote插件数已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[GQA-Plugin] --> sys_role_menu 表中vote插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Save(&sysRoleMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> vote插件初始数据进入 sys_role_menu 表成功！")
		global.GqaLog.Info("[GQA-Plugin] --> vote插件初始数据进入 sys_role_menu 表成功！")
		return nil
	})
}
