package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDict struct{}

func (a *ApiDict) GetDictList(c *gin.Context) {
	var requestDictList model.RequestGetDictList
	if err := model.RequestShouldBindJSON(c, &requestDictList); err != nil {
		return
	}
	if err, dictList, total, parentCode := servicePrivate.ServiceDict.GetDictList(requestDictList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePageWithParentId{
			Records:    dictList,
			Page:       requestDictList.Page,
			PageSize:   requestDictList.PageSize,
			Total:      total,
			ParentCode: parentCode,
		}, c)
	}
}

func (a *ApiDict) EditDict(c *gin.Context) {
	var toEditDict model.SysDict
	if err := model.RequestShouldBindJSON(c, &toEditDict); err != nil {
		return
	}
	toEditDict.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceDict.EditDict(toEditDict); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiDict) AddDict(c *gin.Context) {
	var toAddDict model.RequestAddDict
	if err := model.RequestShouldBindJSON(c, &toAddDict); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDict.Status,
			Sort:      toAddDict.Sort,
			Memo:      toAddDict.Memo,
		},
	}
	addDict := &model.SysDict{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddDict.ParentCode,
		DictCode:                          toAddDict.DictCode,
		DictLabel:                         toAddDict.DictLabel,
		DictExt1:                          toAddDict.DictExt1,
		DictExt2:                          toAddDict.DictExt2,
		DictExt3:                          toAddDict.DictExt3,
		DictExt4:                          toAddDict.DictExt4,
		DictExt5:                          toAddDict.DictExt5,
	}
	if err := servicePrivate.ServiceDict.AddDict(*addDict); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiDict) DeleteDictById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDict.DeleteDictById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiDict) QueryDictById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dict := servicePrivate.ServiceDict.QueryDictById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dict}, utils.GqaI18n("FindSuccess"), c)
	}
}
