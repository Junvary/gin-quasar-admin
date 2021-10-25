package system

import "gin-quasar-admin/global"

type SysDict struct {
	global.GqaModel
	ParentId uint   `json:"parentId" gorm:"comment:父字典ID;"`
	Value    string `json:"value" gorm:"comment:字典编码;not null;uniqueIndex;"`
	Label    string `json:"label" gorm:"comment:字典名称;not null;"`
}

type RequestAddDict struct {
	ParentId uint   `json:"parentId"`
	Value    string `json:"value"`
	Label    string `json:"label"`
}
