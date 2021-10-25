package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceDictDetail struct {
}

func (s *ServiceDictDetail)GetDictDetailList(pageInfo system.RequestPage) (err error, role interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Model(&system.SysDict{})
	var dictList []system.SysDict
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&dictList).Error
	return err, dictList, total
}