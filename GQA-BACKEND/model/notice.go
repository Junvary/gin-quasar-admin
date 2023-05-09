package model

import (
	"github.com/google/uuid"
)

type SysNotice struct {
	GqaModelWithCreatedByAndUpdatedBy
	NoticeId         uuid.UUID         `json:"notice_id" gorm:"comment:消息ID;index;not null"`
	NoticeTitle      string            `json:"notice_title" gorm:"comment:消息题目;index;"`
	NoticeContent    string            `json:"notice_content" gorm:"comment:消息内容;type:text;"`
	NoticeType       string            `json:"notice_type" gorm:"comment:消息类型;index;"`
	NoticeSent       string            `json:"notice_sent" gorm:"comment:已经发送;default:yesNo_no;index;"`
	NoticeToUserType string            `json:"notice_to_user_type" gorm:"comment:接收用户范围;not null;"`
	NoticeToUser     []SysNoticeToUser `json:"notice_to_user" gorm:"foreignKey:NoticeId;references:NoticeId"`
}

type RequestGetNoticeList struct {
	RequestPageAndSort
	NoticeTitle  string `json:"notice_title"`
	NoticeType   string `json:"notice_type"`
	NoticeRead   string `json:"notice_read"`
	NoticeSent   string `json:"notice_sent"`
	NoticeToUser string `json:"notice_to_user"`
}

type RequestAddNotice struct {
	NoticeTitle      string   `json:"notice_title"`
	NoticeContent    string   `json:"notice_content"`
	NoticeType       string   `json:"notice_type"`
	NoticeToUserType string   `json:"notice_to_user_type"`
	NoticeToUser     []string `json:"notice_to_user"`
}
