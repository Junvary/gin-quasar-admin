package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
)

type ServiceGetDict struct {}

func (s *ServiceGetDict)GetDict() (err error, dict []system.SysDict)  {
	var dictList []system.SysDict
	err = global.GqaDb.Find(&dictList).Error
	return err, dictList
}