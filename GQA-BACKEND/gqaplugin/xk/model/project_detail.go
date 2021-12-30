package model

import (
	"github.com/google/uuid"
)

type GqaPluginXkProjectDetail struct {
	ProjectId   uuid.UUID `json:"projectId" gorm:"comment:项目id;index;not null"`
	ProjectNode string    `json:"projectNode" gorm:"comment:项目节点"`
	Content     string    `json:"content" gorm:"comment:描述;type:text"`
	Attachment  string    `json:"attachment" gorm:"comment:附件;type:text;"`
	StartTime   string    `json:"startTime" gorm:"comment:开始时间"`
	EndTime     string    `json:"endTime" gorm:"comment:结束时间"`
}

type RequestEditProjectDetail struct {
	ProjectId     uuid.UUID                  `json:"projectId"`
	ProjectDetail []GqaPluginXkProjectDetail `json:"projectDetail"`
}
