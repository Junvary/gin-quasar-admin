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
		global.GqaLogger.Error("获取Api列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取Api列表失败，"+err.Error(), c)
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
		global.GqaLogger.Error("编辑Api失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑Api失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑Api成功！")
		model.ResponseSuccessMessage("编辑Api成功！", c)
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
		global.GqaLogger.Error("添加Api失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加Api失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加Api成功！", c)
	}
}

func (a *ApiApi) DeleteApiById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceApi.DeleteApiById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除Api失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除Api失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除Api成功！")
		model.ResponseSuccessMessage("删除Api成功！", c)
	}
}

func (a *ApiApi) QueryApiById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, api := servicePrivate.ServiceApi.QueryApiById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找Api失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找Api失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": api}, "查找Api成功！", c)
	}
}
