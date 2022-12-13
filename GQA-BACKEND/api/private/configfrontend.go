package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfigFrontend struct{}

func (a *ApiConfigFrontend) GetConfigFrontendList(c *gin.Context) {
	var getConfigFrontendList model.RequestGetConfigFrontendList
	if err := model.RequestShouldBindJSON(c, &getConfigFrontendList); err != nil {
		return
	}
	if err, configList, total := servicePrivate.ServiceConfigFrontend.GetConfigFrontendList(getConfigFrontendList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  configList,
			Page:     getConfigFrontendList.Page,
			PageSize: getConfigFrontendList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfigFrontend) EditConfigFrontend(c *gin.Context) {
	var toEditConfigFrontend model.SysConfigFrontend
	if err := model.RequestShouldBindJSON(c, &toEditConfigFrontend); err != nil {
		return
	}
	toEditConfigFrontend.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceConfigFrontend.EditConfigFrontend(toEditConfigFrontend); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiConfigFrontend) AddConfigFrontend(c *gin.Context) {
	var toAddConfigFrontend model.RequestAddConfigFrontend
	if err := model.RequestShouldBindJSON(c, &toAddConfigFrontend); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddConfigFrontend.Status,
			Sort:      toAddConfigFrontend.Sort,
			Memo:      toAddConfigFrontend.Memo,
		},
	}
	addConfigFrontend := &model.SysConfigFrontend{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ConfigItem:                        toAddConfigFrontend.ConfigItem,
		ItemDefault:                       toAddConfigFrontend.ItemDefault,
		ItemCustom:                        toAddConfigFrontend.ItemCustom,
	}
	if err := servicePrivate.ServiceConfigFrontend.AddConfigFrontend(*addConfigFrontend); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiConfigFrontend) DeleteConfigFrontendById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceConfigFrontend.DeleteConfigFrontendById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}
