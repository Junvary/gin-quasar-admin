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

func GetConfigFrontend(gqaOption string) (key string) {
	var sysConfigFrontend system.SysConfigFrontend
	if err := global.GqaDb.Where("gqa_option = ?", gqaOption).First(&sysConfigFrontend).Error; err != nil {
		return ""
	}
	if sysConfigFrontend.Custom != "" {
		return sysConfigFrontend.Custom
	} else {
		return sysConfigFrontend.Default
	}
}

func GetDict(dictCode string) (err error, dict []system.SysDict) {
	if err = global.GqaDb.Where("parent_code = ?", dictCode).Find(&dict).Error; err != nil {
		return err, nil
	}
	return nil, dict
}
