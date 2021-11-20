package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

func GetConfigBackend(gqaOption string) (key string) {
	var sysConfigBackend system.SysConfigBackend
	if err := global.GqaDb.Where("gqa_option = ?", gqaOption).First(&sysConfigBackend).Error; err != nil {
		return ""
	}
	if sysConfigBackend.Custom != "" {
		return sysConfigBackend.Custom
	} else {
		return sysConfigBackend.Default
	}
}
