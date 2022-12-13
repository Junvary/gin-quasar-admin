package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogOperation struct{}

func (a *ApiLogOperation) GetLogOperationList(c *gin.Context) {
	var requestLogList model.RequestGetLogOperationList
	if err := model.RequestShouldBindJSON(c, &requestLogList); err != nil {
		return
	}
	if err, logList, total := servicePrivate.ServiceLogOperation.GetLogOperationList(requestLogList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  logList,
			Page:     requestLogList.Page,
			PageSize: requestLogList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiLogOperation) DeleteLogOperationById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceLogOperation.DeleteLogOperationById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}
