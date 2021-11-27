package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceGetFrontend struct {}

func (s *ServiceGetFrontend) GetFrontend() (err error, frontend []system.SysConfigFrontend) {
	var configFrontend []system.SysConfigFrontend
	err = global.GqaDb.Find(&configFrontend).Error
	return err, configFrontend
}
