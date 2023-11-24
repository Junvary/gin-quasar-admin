package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiConfigBackend struct{}

func (a *ApiConfigBackend) GetConfigBackendList(c *gin.Context) {
	var toGetDataList model.RequestGetConfigBackendList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceConfigBackend.GetConfigBackendList(toGetDataList); err != nil {
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

func (a *ApiConfigBackend) EditConfigBackend(c *gin.Context) {
	var toEditData model.SysConfigBackend
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceConfigBackend.EditConfigBackend(toEditData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditSuccess"), c)
	}
}

func (a *ApiConfigBackend) AddConfigBackend(c *gin.Context) {
	var toAddData model.RequestAddConfigBackend
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
	addData := &model.SysConfigBackend{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ConfigItem:                        toAddData.ConfigItem,
		ItemDefault:                       toAddData.ItemDefault,
		ItemCustom:                        toAddData.ItemCustom,
	}
	if err := servicePrivate.ServiceConfigBackend.AddConfigBackend(c, *addData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiConfigBackend) DeleteConfigBackendById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceConfigBackend.DeleteConfigBackendById(c, toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}
