package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {
}

func (a *ApiApi) GetApiList(c *gin.Context) {
	var pageInfo global.RequestPage
	if err := c.ShouldBindJSON(&pageInfo); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败！"+err.Error(), c)
		return
	}
	if err, apiList, total := service.GroupServiceApp.ServiceSystem.GetApiList(pageInfo); err != nil {
		global.GqaLog.Error("获取API列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取API列表失败，" + err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:     apiList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiApi) EditApi(c *gin.Context) {
	var toEditApi system.SysApi
	if err := c.ShouldBindJSON(&toEditApi); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditApi(toEditApi); err != nil {
		global.GqaLog.Error("编辑API失败！", zap.Any("err", err))
		global.ErrorMessage("编辑API失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("编辑API成功！", c)
	}
}

func (a *ApiApi) AddApi(c *gin.Context) {
	var toAddApi system.RequestAddApi
	if err := c.ShouldBindJSON(&toAddApi); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addApi := &system.SysApi{
		GqaModel: global.GqaModel{
			Status: toAddApi.Status,
			Sort:   toAddApi.Sort,
			Remark:   toAddApi.Remark,
		},
		Group: toAddApi.Group,
		Path: toAddApi.Path,
		Method: toAddApi.Method,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddApi(*addApi); err != nil {
		global.GqaLog.Error("添加API失败！", zap.Any("err", err))
		global.ErrorMessage("添加API失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加API成功！", c)
	}
}

func (a *ApiApi) DeleteApi(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.DeleteApi(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除API失败！", zap.Any("err", err))
		global.ErrorMessage("删除API失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("删除API成功！", c)
	}
}

