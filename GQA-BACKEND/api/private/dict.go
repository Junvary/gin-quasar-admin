package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiDict struct{}

func (a *ApiDict) GetDictList(c *gin.Context) {
	var toGetDataList model.RequestGetDictList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total, parentCode := servicePrivate.ServiceDict.GetDictList(toGetDataList); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePageWithParentId{
			Records:    dataList,
			Page:       toGetDataList.Page,
			PageSize:   toGetDataList.PageSize,
			Total:      total,
			ParentCode: parentCode,
		}, c)
	}
}

func (a *ApiDict) EditDict(c *gin.Context) {
	var toEditData model.SysDict
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceDict.EditDict(c, toEditData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditSuccess"), c)
	}
}

func (a *ApiDict) AddDict(c *gin.Context) {
	var toAddData model.RequestAddDict
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
	addData := &model.SysDict{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddData.ParentCode,
		DictCode:                          toAddData.DictCode,
		DictLabel:                         toAddData.DictLabel,
		DictExt1:                          toAddData.DictExt1,
		DictExt2:                          toAddData.DictExt2,
		DictExt3:                          toAddData.DictExt3,
		DictExt4:                          toAddData.DictExt4,
		DictExt5:                          toAddData.DictExt5,
	}
	if err := servicePrivate.ServiceDict.AddDict(c, *addData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiDict) DeleteDictById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDict.DeleteDictById(c, toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiDict) QueryDictById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceDict.QueryDictById(toQueryId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}
