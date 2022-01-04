package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/google/uuid"
)

type GqaPluginXkProject struct {
	UpdatedByUser *system.SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	ProjectId      uuid.UUID                  `json:"projectId" gorm:"comment:项目ID;index;not null"`
	ProjectName    string                     `json:"projectName" gorm:"comment:项目名称;not null;index"`
	Demand         string                     `json:"demand" gorm:"comment:需求单位;"`
	LeaderUsername string                     `json:"leaderUsername" gorm:"comment:牵头人username;"`
	Leader         system.SysUser             `json:"leader" gorm:"comment:牵头人;foreignKey:LeaderUsername;references:Username"`
	Player         string                     `json:"player" gorm:"comment:参与人;"`
	Language       string                     `json:"language" gorm:"comment:项目语言"`
	Node           string                     `json:"node" gorm:"comment:项目节点"`
	ProjectDetail  []GqaPluginXkProjectDetail `json:"projectDetail" gorm:"foreignKey:ProjectId;references:ProjectId"`
}

type RequestAddProject struct {
	Status         string                     `json:"status"`
	Sort           uint                       `json:"sort"`
	Remark         string                     `json:"remark"`
	ProjectName    string                     `json:"projectName"`
	Demand         string                     `json:"demand"`
	LeaderUsername string                     `json:"leaderUsername"`
	Player         string                     `json:"player"`
	Language       string                     `json:"language"`
	Node           string                     `json:"node"`
	ProjectDetail  []GqaPluginXkProjectDetail `json:"projectDetail"`
}

type RequestProjectList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddProject 中的字段
	ProjectName string `json:"projectName"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginXkProject
}
