package system

import "gin-quasar-admin/global"

type SysDict struct {
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
