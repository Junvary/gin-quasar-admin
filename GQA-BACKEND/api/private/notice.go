package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiNotice struct{}

func (a *ApiNotice) GetNoticeList(c *gin.Context) {
	var toGetDataList model.RequestGetNoticeList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceNotice.GetNoticeList(toGetDataList); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
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
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiNotice) DeleteNoticeById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceNotice.DeleteNoticeById(toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiNotice) QueryNoticeById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceNotice.QueryNoticeById(toQueryId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}

func (a *ApiNotice) QueryNoticeReadById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err, data := servicePrivate.ServiceNotice.QueryNoticeReadById(toQueryId.Id, username); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}

func (a *ApiNotice) SendNotice(c *gin.Context) {
	var toSendNotice model.SysNotice
	if err := model.RequestShouldBindJSON(c, &toSendNotice); err != nil {
		return
	}
	toSendNotice.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.SendNotice(toSendNotice); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "SendFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "SendSuccess"), c)
	}
}
