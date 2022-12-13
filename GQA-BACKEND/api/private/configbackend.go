package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfigBackend struct{}

func (a *ApiConfigBackend) GetConfigBackendList(c *gin.Context) {
	var getConfigBackendList model.RequestGetConfigBackendList
	if err := model.RequestShouldBindJSON(c, &getConfigBackendList); err != nil {
		return
	}
	if err, configList, total := servicePrivate.ServiceConfigBackend.GetConfigBackendList(getConfigBackendList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  configList,
			Page:     getConfigBackendList.Page,
			PageSize: getConfigBackendList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfigBackend) EditConfigBackend(c *gin.Context) {
	var toEditConfigBackend model.SysConfigBackend
	if err := model.RequestShouldBindJSON(c, &toEditConfigBackend); err != nil {
		return
	}
	toEditConfigBackend.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceConfigBackend.EditConfigBackend(toEditConfigBackend); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiConfigBackend) AddConfigBackend(c *gin.Context) {
	var toAddConfigBackend model.RequestAddConfigBackend
	if err := model.RequestShouldBindJSON(c, &toAddConfigBackend); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddConfigBackend.Status,
			Sort:      toAddConfigBackend.Sort,
			Memo:      toAddConfigBackend.Memo,
		},
	}
	addConfigBackend := &model.SysConfigBackend{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ConfigItem:                        toAddConfigBackend.ConfigItem,
		ItemDefault:                       toAddConfigBackend.ItemDefault,
		ItemCustom:                        toAddConfigBackend.ItemCustom,
	}
	if err := servicePrivate.ServiceConfigBackend.AddConfigBackend(*addConfigBackend); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiConfigBackend) DeleteConfigBackendById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceConfigBackend.DeleteConfigBackendById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}
