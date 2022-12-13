package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct{}

func (a *ApiApi) GetApiList(c *gin.Context) {
	var getApiList model.RequestGetApiList
	if err := model.RequestShouldBindJSON(c, &getApiList); err != nil {
		return
	}
	if err, configList, total := servicePrivate.ServiceApi.GetApiList(getApiList); err != nil {
		global.GqaLogger.Error(utils.GqaI18nWithData("GetSomeListFailed", "Api"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18nWithData("GetSomeListFailed", "Api")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  configList,
			Page:     getApiList.Page,
			PageSize: getApiList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiApi) EditApi(c *gin.Context) {
	var toEditApi model.SysApi
	if err := model.RequestShouldBindJSON(c, &toEditApi); err != nil {
		return
	}
	toEditApi.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceApi.EditApi(toEditApi); err != nil {
		global.GqaLogger.Error(utils.GqaI18nWithData("EditSomeFailed", "Api"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18nWithData("EditSomeFailed", "Api")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18nWithData("EditSomeSuccess", "Api"))
		model.ResponseSuccessMessage(utils.GqaI18nWithData("EditSomeSuccess", "Api"), c)
	}
}

func (a *ApiApi) AddApi(c *gin.Context) {
	var toAddApi model.RequestAddApi
	if err := model.RequestShouldBindJSON(c, &toAddApi); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddApi.Status,
			Sort:      toAddApi.Sort,
			Memo:      toAddApi.Memo,
		},
	}
	addApi := &model.SysApi{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ApiGroup:                          toAddApi.ApiGroup,
		ApiMethod:                         toAddApi.ApiMethod,
		ApiPath:                           toAddApi.ApiPath,
	}
	if err := servicePrivate.ServiceApi.AddApi(*addApi); err != nil {
		global.GqaLogger.Error(utils.GqaI18nWithData("AddSomeFailed", "Api"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18nWithData("AddSomeFailed", "Api")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18nWithData("AddSomeSuccess", "Api"), c)
	}
}

func (a *ApiApi) DeleteApiById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceApi.DeleteApiById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18nWithData("DeleteSomeFailed", "Api"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18nWithData("DeleteSomeFailed", "Api")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18nWithData("DeleteSomeSuccess", "Api"))
		model.ResponseSuccessMessage(utils.GqaI18nWithData("DeleteSomeSuccess", "Api"), c)
	}
}

func (a *ApiApi) QueryApiById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, api := servicePrivate.ServiceApi.QueryApiById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18nWithData("FindSomeFailed", "Api"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18nWithData("FindSomeFailed", "Api")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": api}, utils.GqaI18nWithData("FindSomeSuccess", "Api"), c)
	}
}
