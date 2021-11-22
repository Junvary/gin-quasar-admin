package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysDict struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	//ParentCode 对应 DictCode
	ParentCode string `json:"parentCode" gorm:"comment:父字典DictCode;index"`
	DictCode   string `json:"dictCode" gorm:"comment:字典编码;not null;index;"`
	DictLabel  string `json:"dictLabel" gorm:"comment:字典名称;not null;"`
}

type RequestAddDict struct {
	Status     string `json:"status"`
	Sort       uint   `json:"sort"`
	Remark     string `json:"remark"`
	ParentCode string `json:"parentCode"`
	DictCode   string `json:"dictCode"`
	DictLabel  string `json:"dictLabel"`
}

type RequestDictList struct {
	global.RequestPageAndSort
	ParentCode string `json:"parentCode"`
	//可扩充的模糊搜索项，参考上面 RequestAddDict 中的字段
	DictCode  string `json:"dictCode"`
	DictLabel string `json:"dictLabel"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysDict
}
