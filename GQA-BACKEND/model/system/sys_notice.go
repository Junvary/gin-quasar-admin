package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/google/uuid"
)

type SysNotice struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	NoticeId         uuid.UUID         `json:"noticeId" gorm:"comment:消息ID;index;not null"`
	NoticeTitle      string            `json:"noticeTitle" gorm:"comment:消息题目;index;"`
	NoticeContent    string            `json:"noticeContent" gorm:"comment:消息内容;type:text;"`
	NoticeType       string            `json:"noticeType" gorm:"comment:消息类型;index;"`
	NoticeSent       string            `json:"noticeSent" gorm:"comment:已经发送;default:no;index;"`
	NoticeToUserType string            `json:"noticeToUserType" gorm:"comment:接收用户范围;not null;"`
	NoticeToUser     []SysNoticeToUser `json:"noticeToUser" gorm:"foreignKey:NoticeId;references:NoticeId"`
}

type RequestAddNotice struct {
	NoticeTitle      string   `json:"noticeTitle"`
	NoticeContent    string   `json:"noticeContent"`
	NoticeType       string   `json:"noticeType"`
	NoticeToUserType string   `json:"noticeToUserType"`
	NoticeToUser     []string `json:"noticeToUser"`
}

type RequestNoticeList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddNotice 中的字段
	NoticeTitle string `json:"noticeTitle"`
	NoticeType  string `json:"noticeType"`
	NoticeRead  string `json:"noticeRead"`
	NoticeSent  string `json:"noticeSent"`
	//下面这个是前端上面通知栏调用时候的搜索内容
	NoticeToUser string `json:"noticeToUser"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysNotice
}
