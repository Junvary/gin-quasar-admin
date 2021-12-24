package system

import "github.com/google/uuid"

type SysNoticeToUser struct {
	NoticeId uuid.UUID `json:"noticeId" gorm:"comment:消息ID;index;not null;"`
	ToUser   string    `json:"toUser" gorm:"comment:接收用户;index;"`
	UserRead string    `json:"userRead" gorm:"comment:是否阅读;default:no;index;"`
}
