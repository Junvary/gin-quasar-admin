package model

import "github.com/google/uuid"

type SysNoticeToUser struct {
	NoticeId uuid.UUID `json:"notice_id" gorm:"comment:消息ID;index;not null;"`
	ToUser   string    `json:"to_user" gorm:"comment:接收用户;index;"`
	UserRead string    `json:"user_read" gorm:"comment:是否阅读;default:yesNo_no;index;"`
}
