package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceDict struct{}

func (s *ServiceDict) GetDictAll() (err error, dict []model.SysDict) {
	var dictList []model.SysDict
	err = global.GqaDb.Find(&dictList).Error
	return err, dictList
}
