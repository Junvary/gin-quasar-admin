package model

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type GqaPluginXkDownload struct {
	UpdatedByUser *system.SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Title      string `json:"title" gorm:"comment:资源标题;not null;index"`
	Content    string `json:"content" gorm:"comment:资源内容;not null;type:text;"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddDownload struct {
	Status     string `json:"status"`
	Sort       uint   `json:"sort"`
	Remark     string `json:"remark"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type RequestDownloadList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddDownload 中的字段
	Title string `json:"title"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginXkDownload
}
