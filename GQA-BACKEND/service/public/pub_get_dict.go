package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceGetDict struct {}

func (s *ServiceGetDict)GetDict() (err error, dict []system.SysDict)  {
	var dictList []system.SysDict
	err = global.GqaDb.Find(&dictList).Error
	return err, dictList
}