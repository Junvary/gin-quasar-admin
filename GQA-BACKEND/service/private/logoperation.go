package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceLogOperation struct{}

func (s *ServiceLogOperation) GetLogOperationList(requestLogList model.RequestGetLogOperationList) (err error, log interface{}, total int64) {
	pageSize := requestLogList.PageSize
	offset := requestLogList.PageSize * (requestLogList.Page - 1)
	db := global.GqaDb.Model(&model.SysLogOperation{})
	var logList []model.SysLogOperation
	// Search
	if requestLogList.OperationUsername != "" {
		db = db.Where("operation_username like ?", "%"+requestLogList.OperationUsername+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestLogList.SortBy, requestLogList.Desc)).Find(&logList).Error
	return err, logList, total
}

func (s *ServiceLogOperation) DeleteLogOperationById(id uint) (err error) {
	var sysLog model.SysLogOperation
	if err = global.GqaDb.Where("id = ?", id).First(&sysLog).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysLog).Error
	return err
}
