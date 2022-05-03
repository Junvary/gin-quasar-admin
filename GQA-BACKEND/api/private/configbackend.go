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
		global.GqaLogger.Error("获取后台配置列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取后台配置列表失败，"+err.Error(), c)
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
		global.GqaLogger.Error("编辑后台配置失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑后台配置失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑后台配置成功！")
		model.ResponseSuccessMessage("编辑后台配置成功！", c)
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
		global.GqaLogger.Error("添加后台配置失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加后台配置失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加后台配置成功！", c)
	}
}

func (a *ApiConfigBackend) DeleteConfigBackendById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceConfigBackend.DeleteConfigBackendById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除后台配置失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除后台配置失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除后台配置成功！")
		model.ResponseSuccessMessage("删除后台配置成功！", c)
	}
}
