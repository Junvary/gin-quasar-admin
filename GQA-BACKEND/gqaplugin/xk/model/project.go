package model

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type GqaPluginXkProject struct {
	UpdatedByUser *system.SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	ProjectName string         `json:"projectName" gorm:"comment:项目名称;not null;index"`
	Demand      string         `json:"demand" gorm:"comment:需求单位;"`
	LeaderId    uint           `json:"leaderId" gorm:"comment:牵头人ID;"`
	Leader      system.SysUser `json:"leader" gorm:"comment:牵头人;foreignKey:LeaderId;references:Id"`
	Player      string         `json:"player" gorm:"comment:参与人;"`
	Language    string         `json:"language" gorm:"comment:项目语言"`
	Node        string         `json:"node" gorm:"comment:项目节点"`
}

type RequestAddProject struct {
	Status      string `json:"status"`
	Sort        uint   `json:"sort"`
	Remark      string `json:"remark"`
	ProjectName string `json:"projectName"`
	Demand      string `json:"demand"`
	LeaderId    uint   `json:"leaderId"`
	Player      string `json:"player"`
	Language    string `json:"language"`
	Node        string `json:"node"`
}

type RequestProjectList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddProject 中的字段
	ProjectName string `json:"projectName"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginXkProject
}
