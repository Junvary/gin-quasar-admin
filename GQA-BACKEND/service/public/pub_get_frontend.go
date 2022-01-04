package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type ServiceGetFrontend struct {}

func (s *ServiceGetFrontend) GetFrontend() (err error, frontend []system.SysConfigFrontend) {
	var configFrontend []system.SysConfigFrontend
	err = global.GqaDb.Find(&configFrontend).Error
	return err, configFrontend
}
