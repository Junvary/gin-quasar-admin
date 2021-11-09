package system

import "gin-quasar-admin/global"

type SysConfig struct {
	global.GqaModel
	GqaOption string `json:"gqaOption" gorm:"comment:配置项;not null;index;"`
	Default   string `json:"default" gorm:"comment:默认配置;not null;"`
	Custom    string `json:"custom" gorm:"comment:自定义配置;"`
}

type RequestAddConfig struct {
	Status   string `json:"status"`
	Sort     uint   `json:"sort"`
	Remark   string `json:"remark"`
	GqaOption string `json:"gqaOption"`
	Default   string `json:"default"`
	Custom    string `json:"custom"`
}
