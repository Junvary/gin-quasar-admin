package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogLogin struct{}

func (a *ApiLogLogin) GetLogLoginList(c *gin.Context) {
	var requestLogList model.RequestGetLogLoginList
	if err := model.RequestShouldBindJSON(c, &requestLogList); err != nil {
		return
	}
	if err, logList, total := servicePrivate.ServiceLogLogin.GetLogLoginList(requestLogList); err != nil {
		global.GqaLogger.Error("获取登录日志列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取登录日志列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  logList,
			Page:     requestLogList.Page,
			PageSize: requestLogList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiLogLogin) DeleteLogLoginById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceLogLogin.DeleteLogLoginById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除登录日志失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除登录日志失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("删除登录日志成功！", c)
	}
}
