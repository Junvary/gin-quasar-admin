package private_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetDownloadList(getDownloadList model.RequestDownloadList) (err error, download []model.GqaPluginXkDownload, total int64) {
	pageSize := getDownloadList.PageSize
	offset := getDownloadList.PageSize * (getDownloadList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkDownload{})
	var downloadList []model.GqaPluginXkDownload
	//配置搜索
	if getDownloadList.Title != ""{
		db = db.Where("title like ?", "%" + getDownloadList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getDownloadList.SortBy, getDownloadList.Desc)).Preload("CreatedByUser").Find(&downloadList).Error
	return err, downloadList, total
}

func  EditDownload(toEditDownload model.GqaPluginXkDownload) (err error) {
	var download model.GqaPluginXkDownload
	if err = global.GqaDb.Where("id = ?", toEditDownload.Id).First(&download).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditDownload).Error
	return err
}

func AddDownload(toAddDownload model.GqaPluginXkDownload) (err error) {
	err = global.GqaDb.Create(&toAddDownload).Error
	return err
}

func DeleteDownload(id uint) (err error) {
	var download model.GqaPluginXkDownload
	if err = global.GqaDb.Where("id = ?", id).First(&download).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&download).Error
	return err
}

func QueryDownloadById(id uint) (err error, downloadInfo model.GqaPluginXkDownload) {
	var download model.GqaPluginXkDownload
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&download, "id = ?", id).Error
	return err, download
}