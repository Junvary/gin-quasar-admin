package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"time"
)

type PluginAchievementObtain struct {
	CategoryCode  string         `json:"category_code" gorm:"comment:成就编码;not null;index"`
	CreatedBy     string         `json:"created_by" gorm:"comment:获得人;not null;index;"`
	CreatedAt     time.Time      `json:"created_at" gorm:"comment:获取时间;not null;"`
	CreatedByUser *model.SysUser `json:"created_by_user" gorm:"foreignKey:CreatedBy;references:Username"`
}

type RequestGetObtainList struct {
	gqaModel.RequestPageAndSort
	CategoryCode string `json:"category_code"`
}

type ObtainFind struct {
	CategoryCode string `json:"category_code"`
	Username     string `json:"username"`
}
