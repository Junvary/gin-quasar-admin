package public_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetNews(getNewsList model.RequestNewsList) (err error, news []model.GqaPluginXkNews, total int64) {
	pageSize := getNewsList.PageSize
	offset := getNewsList.PageSize * (getNewsList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkNews{})
	var newsList []model.GqaPluginXkNews
	//配置搜索
	if getNewsList.Title != ""{
		db = db.Where("title like ?", "%" + getNewsList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getNewsList.SortBy, getNewsList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}
