package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogLogin struct{}

func (a *ApiLogLogin) GetLogLoginList(c *gin.Context) {
	var requestLogList system.RequestLogLoginList
	if err := c.ShouldBindJSON(&requestLogList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, logList, total := ServiceLogLogin.GetLogLoginList(requestLogList); err != nil {
		global.GqaLog.Error("获取登录日志列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取登录日志列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:    logList,
			Page:       requestLogList.Page,
			PageSize:   requestLogList.PageSize,
			Total:      total,
		}, c)
	}
}

func (a *ApiLogLogin) DeleteLogLogin(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceLogLogin.DeleteLogLogin(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除登录日志失败！", zap.Any("err", err))
		global.ErrorMessage("删除登录日志失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除登录日志成功！", c)
	}
}