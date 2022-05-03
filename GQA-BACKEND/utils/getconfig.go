package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

func GetConfigBackend(configItem string) (key string) {
	var sysConfigBackend model.SysConfigBackend
	if err := global.GqaDb.Where("config_item = ?", configItem).First(&sysConfigBackend).Error; err != nil {
		return ""
	}
	if sysConfigBackend.ItemCustom != "" {
		return sysConfigBackend.ItemCustom
	} else {
		return sysConfigBackend.ItemDefault
	}
}

func GetConfigFrontend(configItem string) (key string) {
	var sysConfigFrontend model.SysConfigFrontend
	if err := global.GqaDb.Where("config_item = ?", configItem).First(&sysConfigFrontend).Error; err != nil {
		return ""
	}
	if sysConfigFrontend.ItemCustom != "" {
		return sysConfigFrontend.ItemCustom
	} else {
		return sysConfigFrontend.ItemDefault
	}
}

func GetDict(dictCode string) (err error, dict []model.SysDict) {
	if err = global.GqaDb.Where("parent_code = ?", dictCode).Find(&dict).Error; err != nil {
		return err, nil
	}
	return nil, dict
}
