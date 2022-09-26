package model

import gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"

type PluginAchievementCategory struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Category string `json:"category" gorm:"comment:成就编码;not null;index"`
	Code     string `json:"code" gorm:"comment:小编码;not null;index"`
	Name     string `json:"name" gorm:"comment:成就名称;not null"`
}

type RequestAddCategory struct {
	gqaModel.RequestAdd
	Category string `json:"category"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

type RequestGetCategoryList struct {
	gqaModel.RequestPageAndSort
	Category string `json:"category"`
}
