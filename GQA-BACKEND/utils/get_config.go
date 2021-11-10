package utils

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

func GetConfig(gqaOption string) (key string) {
	var sysConfig system.SysConfig
	if err := global.GqaDb.Where("gqa_option = ?", gqaOption).First(&sysConfig).Error; err != nil {
		return ""
	}
	if sysConfig.Custom != "" {
		return sysConfig.Custom
	} else {
		return sysConfig.Default
	}
}
