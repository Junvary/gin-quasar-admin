package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfigBackend struct {}

func (a *ApiConfigBackend) GetConfigBackendList(c *gin.Context) {
	var getConfigBackendList system.RequestConfigBackendList
	if err := c.ShouldBindJSON(&getConfigBackendList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, configList, total := ServiceConfigBackend.GetConfigBackendList(getConfigBackendList); err != nil {
		global.GqaLog.Error("获取后台配置列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取后台配置列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  configList,
			Page:     getConfigBackendList.Page,
			PageSize: getConfigBackendList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfigBackend) EditConfigBackend(c *gin.Context) {
	var toEditConfigBackend system.SysConfigBackend
	if err := c.ShouldBindJSON(&toEditConfigBackend); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditConfigBackend.UpdatedBy = utils.GetUsername(c)
	if err := ServiceConfigBackend.EditConfigBackend(toEditConfigBackend); err != nil {
		global.GqaLog.Error("编辑后台配置失败！", zap.Any("err", err))
		global.ErrorMessage("编辑后台配置失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "编辑后台配置成功！")
		global.SuccessMessage("编辑后台配置成功！", c)
	}
}

func (a *ApiConfigBackend) AddConfigBackend(c *gin.Context) {
	var toAddConfigBackend system.RequestAddConfigBackend
	if err := c.ShouldBindJSON(&toAddConfigBackend); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addConfigBackend := &system.SysConfigBackend{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddConfigBackend.Status,
			Sort:      toAddConfigBackend.Sort,
			Remark:    toAddConfigBackend.Remark,
		},
		GqaOption: toAddConfigBackend.GqaOption,
		Default:   toAddConfigBackend.Default,
		Custom:    toAddConfigBackend.Custom,
	}
	if err := ServiceConfigBackend.AddConfigBackend(*addConfigBackend); err != nil {
		global.GqaLog.Error("添加后台配置失败！", zap.Any("err", err))
		global.ErrorMessage("添加后台配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加后台配置成功！", c)
	}
}

func (a *ApiConfigBackend) DeleteConfigBackend(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceConfigBackend.DeleteConfigBackend(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除后台配置失败！", zap.Any("err", err))
		global.ErrorMessage("删除后台配置失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "删除后台配置成功！")
		global.SuccessMessage("删除后台配置成功！", c)
	}
}
