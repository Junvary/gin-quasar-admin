package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfig struct {
}

func (a *ApiConfig) GetConfigList(c *gin.Context) {
	var requestConfigList system.RequestConfigList
	if err := c.ShouldBindJSON(&requestConfigList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, configList, total := ServiceConfig.GetConfigList(requestConfigList); err != nil {
		global.GqaLog.Error("获取配置列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取配置列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  configList,
			Page:     requestConfigList.Page,
			PageSize: requestConfigList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfig) EditConfig(c *gin.Context) {
	var toEditConfig system.SysConfig
	if err := c.ShouldBindJSON(&toEditConfig); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditConfig.UpdatedBy = utils.GetUsername(c)
	if err := ServiceConfig.EditConfig(toEditConfig); err != nil {
		global.GqaLog.Error("编辑配置失败！", zap.Any("err", err))
		global.ErrorMessage("编辑配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑配置成功！", c)
	}
}

func (a *ApiConfig) AddConfig(c *gin.Context) {
	var toAddConfig system.RequestAddConfig
	if err := c.ShouldBindJSON(&toAddConfig); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addConfig := &system.SysConfig{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddConfig.Status,
			Sort:      toAddConfig.Sort,
			Remark:    toAddConfig.Remark,
		},
		GqaOption: toAddConfig.GqaOption,
		Default:   toAddConfig.Default,
		Custom:    toAddConfig.Custom,
	}
	if err := ServiceConfig.AddConfig(*addConfig); err != nil {
		global.GqaLog.Error("添加配置失败！", zap.Any("err", err))
		global.ErrorMessage("添加配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加配置成功！", c)
	}
}

func (a *ApiConfig) DeleteConfig(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceConfig.DeleteConfig(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除配置失败！", zap.Any("err", err))
		global.ErrorMessage("删除配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除配置成功！", c)
	}
}
