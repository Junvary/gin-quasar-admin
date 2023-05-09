package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceLogLogin struct{}

func (s *ServiceLogLogin) GetLogLoginList(requestLogList model.RequestGetLogLoginList) (err error, log interface{}, total int64) {
	pageSize := requestLogList.PageSize
	offset := requestLogList.PageSize * (requestLogList.Page - 1)
	db := global.GqaDb.Model(&model.SysLogLogin{})
	var logList []model.SysLogLogin
	// Search
	if requestLogList.LoginUsername != "" {
		db = db.Where("login_username like ?", "%"+requestLogList.LoginUsername+"%")
	}
	if requestLogList.LoginSuccess != "" {
		db = db.Where("login_success = ?", requestLogList.LoginSuccess)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestLogList.SortBy, requestLogList.Desc)).Find(&logList).Error
	return err, logList, total
}

func (s *ServiceLogLogin) DeleteLogLoginById(id uint) (err error) {
	var sysLog model.SysLogLogin
	if err = global.GqaDb.Where("id = ?", id).First(&sysLog).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysLog).Error
	return err
}
