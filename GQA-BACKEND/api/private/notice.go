package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiNotice struct{}

func (a *ApiNotice) GetNoticeList(c *gin.Context) {
	var toGetDataList model.RequestGetNoticeList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceNotice.GetNoticeList(toGetDataList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiNotice) AddNotice(c *gin.Context) {
	var toAddData model.RequestAddNotice
	if err := model.RequestShouldBindJSON(c, &toAddData); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.AddNotice(toAddData, username); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiNotice) DeleteNoticeById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceNotice.DeleteNoticeById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiNotice) QueryNoticeById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceNotice.QueryNoticeById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n("FindSuccess"), c)
	}
}

func (a *ApiNotice) QueryNoticeReadById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err, data := servicePrivate.ServiceNotice.QueryNoticeReadById(toQueryId.Id, username); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n("FindSuccess"), c)
	}
}

func (a *ApiNotice) SendNotice(c *gin.Context) {
	var toSendNotice model.SysNotice
	if err := model.RequestShouldBindJSON(c, &toSendNotice); err != nil {
		return
	}
	toSendNotice.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.SendNotice(toSendNotice); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("SendFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("SendFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("SendSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("SendSuccess"), c)
	}
}
