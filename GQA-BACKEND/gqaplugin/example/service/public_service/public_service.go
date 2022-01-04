package public_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/model"
)

func GetNews() (err error, news []model.GqaPluginExampleNews) {
	var newsList []model.GqaPluginExampleNews
	err = global.GqaDb.Find(&newsList).Error
	return err, newsList
}
