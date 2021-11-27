package public_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/example/model"
)

func GetNews() (err error, news []model.GqaPluginExampleNews) {
	var newsList []model.GqaPluginExampleNews
	err = global.GqaDb.Find(&newsList).Error
	return err, newsList
}
