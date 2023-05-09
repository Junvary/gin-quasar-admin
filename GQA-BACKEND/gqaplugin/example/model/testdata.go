package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type PluginExampleTestData struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	Column1 string `json:"column_1" gorm:"comment:第1列;"`
	Column2 string `json:"column_2" gorm:"comment:第2列;"`
	Column3 string `json:"column_3" gorm:"comment:第3列;"`
	Column4 string `json:"column_4" gorm:"comment:第4列;"`
	Column5 string `json:"column_5" gorm:"comment:第5列;"`
}

type RequestAddTestData struct {
	gqaModel.RequestAdd
	Column1 string `json:"column_1"`
	Column2 string `json:"column_2"`
	Column3 string `json:"column_3"`
	Column4 string `json:"column_4"`
	Column5 string `json:"column_5"`
}

type RequestGetTestDataList struct {
	gqaModel.RequestPageAndSort
	Column1 string `json:"column_1"`
}

type RequestFilename struct {
	Filename string `json:"filename"`
}
