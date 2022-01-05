package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"gorm.io/gorm"
)

func GetDownloadList(getDownloadList model.RequestDownloadList, username string) (err error, download []model.GqaPluginXkDownload, total int64) {
	pageSize := getDownloadList.PageSize
	offset := getDownloadList.PageSize * (getDownloadList.Page - 1)
	var downloadList []model.GqaPluginXkDownload
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkDownload{})); err != nil {
		return err, downloadList, 0
	}
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

func  EditDownload(toEditDownload model.GqaPluginXkDownload, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkDownload{})); err != nil {
		return err
	}
	var download model.GqaPluginXkDownload
	if err = db.Where("id = ?", toEditDownload.Id).First(&download).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditDownload).Error
	err = db.Updates(utils.MergeMap(utils.GlobalModelToMap(&toEditDownload.GqaModel), map[string]interface{}{
		"title":      toEditDownload.Title,
		"content":    toEditDownload.Content,
		"attachment": toEditDownload.Attachment,
	})).Error
	return err
}

func AddDownload(toAddDownload model.GqaPluginXkDownload, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkDownload{})); err != nil {
		return err
	}
	err = db.Create(&toAddDownload).Error
	return err
}

func DeleteDownload(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkDownload{})); err != nil {
		return err
	}
	var download model.GqaPluginXkDownload
	if err = db.Where("id = ?", id).First(&download).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&download).Error
	return err
}

func QueryDownloadById(id uint, username string) (err error, downloadInfo model.GqaPluginXkDownload) {
	var download model.GqaPluginXkDownload
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkDownload{})); err != nil {
		return err, download
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&download, "id = ?", id).Error
	return err, download
}