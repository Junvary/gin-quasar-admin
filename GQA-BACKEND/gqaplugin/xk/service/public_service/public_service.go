package public_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetNews() (err error, news []model.GqaPluginXkNews) {
	var newsList []model.GqaPluginXkNews
	err = global.GqaDb.Find(&newsList).Error
	return err, newsList
}
