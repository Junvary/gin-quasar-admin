package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	adapter "github.com/casbin/gorm-adapter/v3"
)

type ServiceApi struct {
}

func (s *ServiceApi) GetApiList(pageInfo system.RequestPage) (err error, api interface{}, total int64) {
	pageSize := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GqaDb.Where("v0=?", "super-admin").Find(&adapter.CasbinRule{})
	var apiList []adapter.CasbinRule
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&apiList).Error
	return err, apiList, total
}
