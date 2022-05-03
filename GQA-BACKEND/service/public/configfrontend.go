package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceConfigFrontend struct{}

func (s *ServiceConfigFrontend) GetConfigFrontendAll() (err error, frontend []model.SysConfigFrontend) {
	var configList []model.SysConfigFrontend
	err = global.GqaDb.Find(&configList).Error
	return err, configList
}
