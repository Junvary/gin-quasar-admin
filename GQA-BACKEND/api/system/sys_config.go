package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfig struct {
}

func (a *ApiConfig) GetConfigList(c *gin.Context) {
	var pageInfo global.RequestPage
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, configList, total := service.GroupServiceApp.ServiceSystem.GetConfigList(pageInfo); err != nil {
		global.GqaLog.Error("获取配置列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取配置列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  configList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfig) EditConfig(c *gin.Context) {
	var toEditConfig system.SysConfig
	if err := c.ShouldBindJSON(&toEditConfig); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditConfig(toEditConfig); err != nil {
		global.GqaLog.Error("编辑配置失败！", zap.Any("err", err))
		global.ErrorMessage("编辑配置失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("编辑配置成功！", c)
	}
}

func (a *ApiConfig) AddConfig(c *gin.Context) {
	var toAddConfig system.RequestAddConfig
	if err := c.ShouldBindJSON(&toAddConfig); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addConfig := &system.SysConfig{
		GqaModel: global.GqaModel{
			Status: toAddConfig.Status,
			Sort:   toAddConfig.Sort,
			Remark:   toAddConfig.Remark,
		},
		GqaOption: toAddConfig.GqaOption,
		Default: toAddConfig.Default,
		Custom: toAddConfig.Custom,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddConfig(*addConfig); err != nil {
		global.GqaLog.Error("添加配置失败！", zap.Any("err", err))
		global.ErrorMessage("添加配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加配置成功！", c)
	}
}
