package system

import "gin-quasar-admin/global"

type SysNotice struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	NoticeTitle   string `json:"noticeTitle" gorm:"comment:消息题目;index;"`
	NoticeContent string `json:"noticeContent" gorm:"comment:消息内容;type:text;"`
	NoticeType    string `json:"noticeType" gorm:"comment:消息类型;index;"`
	NoticeRead    string `json:"noticeRead" gorm:"comment:是否阅读;default:no;index;"`
	NoticeSent    string `json:"noticeSent" gorm:"comment:已经发送;default:no;index;"`
	NoticeToUser  string `json:"noticeToUser" gorm:"comment:接收用户;not null;"`
}

type RequestAddNotice struct {
	NoticeTitle   string `json:"noticeTitle"`
	NoticeContent string `json:"noticeContent"`
	NoticeType    string `json:"noticeType"`
	NoticeToUser  string `json:"noticeToUser"`
}

type RequestNoticeList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddNotice 中的字段
	NoticeTitle string `json:"noticeTitle"`
	NoticeType  string `json:"noticeType"`
	NoticeRead  string `json:"noticeRead"`
	NoticeSent  string `json:"noticeSent"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysNotice
}
