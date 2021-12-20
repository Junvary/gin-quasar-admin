package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogOperation struct{}

func (a *ApiLogOperation) GetLogOperationList(c *gin.Context) {
	var requestLogList system.RequestLogOperationList
	if err := c.ShouldBindJSON(&requestLogList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, logList, total := ServiceLogOperation.GetLogOperationList(requestLogList); err != nil {
		global.GqaLog.Error("获取操作日志列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取操作日志列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:    logList,
			Page:       requestLogList.Page,
			PageSize:   requestLogList.PageSize,
			Total:      total,
		}, c)
	}
}

func (a *ApiLogOperation) DeleteLogOperation(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceLogOperation.DeleteLogOperation(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除操作日志失败！", zap.Any("err", err))
		global.ErrorMessage("删除操作日志失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除操作日志成功！", c)
	}
}