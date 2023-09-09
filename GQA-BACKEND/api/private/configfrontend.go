package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiConfigFrontend struct{}

func (a *ApiConfigFrontend) GetConfigFrontendList(c *gin.Context) {
	var toGetDataList model.RequestGetConfigFrontendList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceConfigFrontend.GetConfigFrontendList(toGetDataList); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "GetListFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfigFrontend) EditConfigFrontend(c *gin.Context) {
	var toEditData model.SysConfigFrontend
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceConfigFrontend.EditConfigFrontend(toEditData); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "EditFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		global.GqaSLogger.Info(utils.GetUsername(c) + utils.GqaI18n(c, "EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n(c, "EditSuccess"), c)
	}
}

func (a *ApiConfigFrontend) AddConfigFrontend(c *gin.Context) {
	var toAddData model.RequestAddConfigFrontend
	if err := model.RequestShouldBindJSON(c, &toAddData); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddData.Status,
			Sort:      toAddData.Sort,
			Memo:      toAddData.Memo,
		},
	}
	addData := &model.SysConfigFrontend{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ConfigItem:                        toAddData.ConfigItem,
		ItemDefault:                       toAddData.ItemDefault,
		ItemCustom:                        toAddData.ItemCustom,
	}
	if err := servicePrivate.ServiceConfigFrontend.AddConfigFrontend(*addData); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "AddFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiConfigFrontend) DeleteConfigFrontendById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceConfigFrontend.DeleteConfigFrontendById(toDeleteId.Id); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "DeleteFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		global.GqaSLogger.Info(utils.GetUsername(c) + utils.GqaI18n(c, "DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}
