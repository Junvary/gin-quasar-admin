package public_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqa_plugin/xk/model"
)

func GetNews() (err error, news []model.GqaPluginXkNews) {
	var newsList []model.GqaPluginXkNews
	err = global.GqaDb.Model(&newsList).Error
	return err, newsList
}
