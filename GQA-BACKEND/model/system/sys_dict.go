package system

import "gin-quasar-admin/global"

type SysDict struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	ParentId uint   `json:"parentId" gorm:"comment:父字典ID;not null;index"`
	Value    string `json:"value" gorm:"comment:字典编码;not null;index;"`
	Label    string `json:"label" gorm:"comment:字典名称;not null;"`
}

type RequestAddDict struct {
	Status   string `json:"status"`
	Sort     uint   `json:"sort"`
	Remark   string `json:"remark"`
	ParentId uint   `json:"parentId"`
	Value    string `json:"value"`
	Label    string `json:"label"`
}

type RequestDictList struct {
	global.RequestPageAndSort
	ParentId uint   `json:"parentId"`
	//可扩充的模糊搜索项，参考上面 RequestAddDict 中的字段
	Value    string `json:"value"`
	Label    string `json:"label"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysDict
}
