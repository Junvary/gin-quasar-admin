package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type ServiceGetBackend struct{}

func (s *ServiceGetBackend) GetBackend() (err error, backend []system.SysConfigBackend) {
	var configBackend []system.SysConfigBackend
	err = global.GqaDb.
		Where("gqa_option = ?", "avatarMaxSize").
		Or("gqa_option = ?", "avatarExt").
		Or("gqa_option = ?", "fileMaxSize").
		Or("gqa_option = ?", "fileExt").
		Or("gqa_option = ?", "webLogoMaxSize").
		Or("gqa_option = ?", "webLogoExt").
		Or("gqa_option = ?", "headerLogoMaxSize").
		Or("gqa_option = ?", "headerLogoExt").Find(&configBackend).Error
	return err, configBackend
}
