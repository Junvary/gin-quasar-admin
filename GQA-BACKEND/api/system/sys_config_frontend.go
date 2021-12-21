package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfigFrontend struct {}

func (a *ApiConfigFrontend) GetConfigFrontendList(c *gin.Context) {
	var getConfigFrontendList system.RequestConfigFrontendList
	if err := c.ShouldBindJSON(&getConfigFrontendList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, configList, total := ServiceConfigFrontend.GetConfigFrontendList(getConfigFrontendList); err != nil {
		global.GqaLog.Error("获取网站前台配置列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取网站前台配置列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  configList,
			Page:     getConfigFrontendList.Page,
			PageSize: getConfigFrontendList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiConfigFrontend) EditConfigFrontend(c *gin.Context) {
	var toEditConfigFrontend system.SysConfigFrontend
	if err := c.ShouldBindJSON(&toEditConfigFrontend); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditConfigFrontend.UpdatedBy = utils.GetUsername(c)
	if err := ServiceConfigFrontend.EditConfigFrontend(toEditConfigFrontend); err != nil {
		global.GqaLog.Error("编辑网站前台配置失败！", zap.Any("err", err))
		global.ErrorMessage("编辑网站前台配置失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "编辑网站前台配置成功！")
		global.SuccessMessage("编辑网站前台配置成功！", c)
	}
}

func (a *ApiConfigFrontend) AddConfigFrontend(c *gin.Context) {
	var toAddConfigFrontend system.RequestAddConfigFrontend
	if err := c.ShouldBindJSON(&toAddConfigFrontend); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addConfigFrontend := &system.SysConfigFrontend{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddConfigFrontend.Status,
			Sort:      toAddConfigFrontend.Sort,
			Remark:    toAddConfigFrontend.Remark,
		},
		GqaOption: toAddConfigFrontend.GqaOption,
		Default:   toAddConfigFrontend.Default,
		Custom:    toAddConfigFrontend.Custom,
	}
	if err := ServiceConfigFrontend.AddConfigFrontend(*addConfigFrontend); err != nil {
		global.GqaLog.Error("添加网站前台配置失败！", zap.Any("err", err))
		global.ErrorMessage("添加网站前台配置失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加网站前台配置成功！", c)
	}
}

func (a *ApiConfigFrontend) DeleteConfigFrontend(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceConfigFrontend.DeleteConfigFrontend(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除网站前台配置失败！", zap.Any("err", err))
		global.ErrorMessage("删除网站前台配置失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "删除网站前台配置成功！")
		global.SuccessMessage("删除网站前台配置成功！", c)
	}
}
