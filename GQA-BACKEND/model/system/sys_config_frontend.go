package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysConfigFrontend struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	GqaOption string `json:"gqaOption" gorm:"comment:配置项;not null;index;"`
	Default   string `json:"default" gorm:"comment:默认配置;not null;"`
	Custom    string `json:"custom" gorm:"comment:自定义配置;"`
}

type RequestAddConfigFrontend struct {
	Status    string `json:"status"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	GqaOption string `json:"gqaOption"`
	Default   string `json:"default"`
	Custom    string `json:"custom"`
}

type RequestConfigFrontendList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddConfigFrontend 中的字段
	GqaOption string `json:"gqaOption"`
	Remark    string `json:"remark"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysConfigFrontend
}
