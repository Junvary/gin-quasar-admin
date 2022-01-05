package public_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gqa-plugin-xk/model"
)

func GetDownload(getDownloadList model.RequestDownloadList) (err error, news []model.GqaPluginXkDownload, total int64) {
	pageSize := getDownloadList.PageSize
	offset := getDownloadList.PageSize * (getDownloadList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkDownload{})
	var newsList []model.GqaPluginXkDownload
	//配置搜索
	if getDownloadList.Title != ""{
		db = db.Where("title like ?", "%" + getDownloadList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getDownloadList.SortBy, getDownloadList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
