package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type GqaPluginXkNews struct {
	UpdatedByUser *system.SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Title      string `json:"title" gorm:"comment:新闻标题;not null;index"`
	Content    string `json:"content" gorm:"comment:新闻内容;not null;type:text;"`
	Attachment string `json:"attachment" gorm:"comment:附件;type:text;"`
}

type RequestAddNews struct {
	Status     string `json:"status"`
	Sort       uint   `json:"sort"`
	Remark     string `json:"remark"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type RequestNewsList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddNews 中的字段
	Title string `json:"title"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginXkNews
}
