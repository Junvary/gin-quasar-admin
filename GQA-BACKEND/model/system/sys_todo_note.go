package system

import "gin-quasar-admin/global"

type SysTodoNote struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	TodoDetail string `json:"todoDetail" gorm:"comment:内容;type:text;"`
	TodoStatus string `json:"todoStatus" gorm:"comment:状态;default:no;index;"`
}

type RequestAddTodoNote struct {
	TodoDetail string `json:"todoDetail"`
}

type RequestTodoNoteList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddTodoNote 中的字段
	TodoStatus string `json:"todoStatus"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysTodoNote
}
