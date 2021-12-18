package private_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
	"gin-quasar-admin/service/system"
	"gorm.io/gorm"
)

func GetNewsList(getNewsList model.RequestNewsList, username string) (err error, news []model.GqaPluginXkNews, total int64) {
	pageSize := getNewsList.PageSize
	offset := getNewsList.PageSize * (getNewsList.Page - 1)
	//db := global.GqaDb.Model(&model.GqaPluginXkNews{})
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err!=nil{
		return
	}
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

func  EditNews(toEditNews model.GqaPluginXkNews) (err error) {
	var news model.GqaPluginXkNews
	if err = global.GqaDb.Where("id = ?", toEditNews.Id).First(&news).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditNews).Error
	return err
}

func AddNews(toAddNews model.GqaPluginXkNews) (err error) {
	err = global.GqaDb.Create(&toAddNews).Error
	return err
}

func DeleteNews(id uint) (err error) {
	var news model.GqaPluginXkNews
	if err = global.GqaDb.Where("id = ?", id).First(&news).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&news).Error
	return err
}

func QueryNewsById(id uint) (err error, newsInfo model.GqaPluginXkNews) {
	var news model.GqaPluginXkNews
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&news, "id = ?", id).Error
	return err, news
}