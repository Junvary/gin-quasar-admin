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
		global.GqaLogger.Error("获取字典列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取字典列表失败，"+err.Error(), c)
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
		global.GqaLogger.Error("编辑字典失败!", zap.Any("err", err))
		model.ResponseErrorMessage("编辑字典失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑字典成功！")
		model.ResponseSuccessMessage("编辑字典成功！", c)
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
	}
	if err := servicePrivate.ServiceDict.AddDict(*addDict); err != nil {
		global.GqaLogger.Error("添加字典失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加字典失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加字典成功！", c)
	}
}

func (a *ApiDict) DeleteDictById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDict.DeleteDictById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除字典失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除字典失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除字典成功！")
		model.ResponseSuccessMessage("删除字典成功！", c)
	}
}

func (a *ApiDict) QueryDictById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dict := servicePrivate.ServiceDict.QueryDictById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找字典失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找字典失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dict}, "查找字典成功！", c)
	}
}
