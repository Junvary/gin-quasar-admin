package public_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gqa-plugin-xk/model"
)

func GetDocument(getDocumentList model.RequestDocumentList) (err error, news []model.GqaPluginXkDocument, total int64) {
	pageSize := getDocumentList.PageSize
	offset := getDocumentList.PageSize * (getDocumentList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkDocument{})
	var newsList []model.GqaPluginXkDocument
	//配置搜索
	if getDocumentList.Title != ""{
		db = db.Where("title like ?", "%" + getDocumentList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getDocumentList.SortBy, getDocumentList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
