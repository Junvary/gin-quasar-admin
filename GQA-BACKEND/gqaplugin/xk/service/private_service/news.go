package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"gorm.io/gorm"
)

func GetNewsList(getNewsList model.RequestNewsList, username string) (err error, news []model.GqaPluginXkNews, total int64) {
	pageSize := getNewsList.PageSize
	offset := getNewsList.PageSize * (getNewsList.Page - 1)
	var newsList []model.GqaPluginXkNews
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err != nil {
		return err, newsList, 0
	}
	//配置搜索
	if getNewsList.Title != "" {
		db = db.Where("title like ?", "%"+getNewsList.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, newsList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getNewsList.SortBy, getNewsList.Desc)).Preload("CreatedByUser").Find(&newsList).Error
	return err, newsList, total
}

func EditNews(toEditNews model.GqaPluginXkNews, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err != nil {
		return err
	}
	var news model.GqaPluginXkNews
	if err = db.Where("id = ?", toEditNews.Id).First(&news).Error; err != nil {
		return err
	}
	//err = global.GqaDb.Updates(&toEditNews).Error
	err = db.Updates(utils.MergeMap(utils.GlobalModelToMap(&toEditNews.GqaModel), map[string]interface{}{
		"title":      toEditNews.Title,
		"content":    toEditNews.Content,
		"attachment": toEditNews.Attachment,
	})).Error
	return err
}

func AddNews(toAddNews model.GqaPluginXkNews, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err != nil {
		return err
	}
	err = db.Create(&toAddNews).Error
	return err
}

func DeleteNews(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err != nil {
		return err
	}
	var news model.GqaPluginXkNews
	if err = db.Where("id = ?", id).First(&news).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&news).Error
	return err
}

func QueryNewsById(id uint, username string) (err error, newsInfo model.GqaPluginXkNews) {
	var news model.GqaPluginXkNews
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkNews{})); err != nil {
		return err, news
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&news, "id = ?", id).Error
	return err, news
}
