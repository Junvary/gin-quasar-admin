package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceConfigBackend struct{}

func (s *ServiceConfigBackend) GetConfigBackendAll() (err error, backend []model.SysConfigBackend) {
	var configList []model.SysConfigBackend
	configItem := []string{
		"avatarMaxSize", "avatarExt",
		"fileMaxSize", "fileExt",
		"logoMaxSize", "logoExt",
		"faviconMaxSize", "faviconExt",
		"bannerImageMaxSize", "bannerImageExt",
	}
	err = global.GqaDb.Where("config_item in ?", configItem).Find(&configList).Error
	return err, configList
}
