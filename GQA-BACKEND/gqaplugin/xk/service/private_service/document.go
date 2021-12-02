package private_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetDocumentList(getDocumentList model.RequestDocumentList) (err error, document []model.GqaPluginXkDocument, total int64) {
	pageSize := getDocumentList.PageSize
	offset := getDocumentList.PageSize * (getDocumentList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkDocument{})
	var documentList []model.GqaPluginXkDocument
	//配置搜索
	if getDocumentList.Title != ""{
		db = db.Where("title like ?", "%" + getDocumentList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getDocumentList.SortBy, getDocumentList.Desc)).Preload("CreatedByUser").Find(&documentList).Error
	return err, documentList, total
}

func  EditDocument(toEditDocument model.GqaPluginXkDocument) (err error) {
	var document model.GqaPluginXkDocument
	if err = global.GqaDb.Where("id = ?", toEditDocument.Id).First(&document).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditDocument).Error
	return err
}

func AddDocument(toAddDocument model.GqaPluginXkDocument) (err error) {
	err = global.GqaDb.Create(&toAddDocument).Error
	return err
}

func DeleteDocument(id uint) (err error) {
	var document model.GqaPluginXkDocument
	if err = global.GqaDb.Where("id = ?", id).First(&document).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&document).Error
	return err
}

func QueryDocumentById(id uint) (err error, documentInfo model.GqaPluginXkDocument) {
	var document model.GqaPluginXkDocument
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&document, "id = ?", id).Error
	return err, document
}