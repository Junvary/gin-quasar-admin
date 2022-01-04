package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {}

func (a *ApiApi) GetApiList(c *gin.Context) {
	var requestApiList system.RequestApiList
	if err := c.ShouldBindJSON(&requestApiList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败！"+err.Error(), c)
		return
	}
	if err, apiList, total := ServiceApi.GetApiList(requestApiList); err != nil {
		global.GqaLog.Error("获取API列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取API列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  apiList,
			Page:     requestApiList.Page,
			PageSize: requestApiList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiApi) EditApi(c *gin.Context) {
	var toEditApi system.SysApi
	if err := c.ShouldBindJSON(&toEditApi); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditApi.UpdatedBy = utils.GetUsername(c)
	if err := ServiceApi.EditApi(toEditApi); err != nil {
		global.GqaLog.Error("编辑API失败！", zap.Any("err", err))
		global.ErrorMessage("编辑API失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "编辑API成功！")
		global.SuccessMessage("编辑API成功！", c)
	}
}

func (a *ApiApi) AddApi(c *gin.Context) {
	var toAddApi system.RequestAddApi
	if err := c.ShouldBindJSON(&toAddApi); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addApi := &system.SysApi{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddApi.Status,
			Sort:      toAddApi.Sort,
			Remark:    toAddApi.Remark,
		},
		ApiGroup:  toAddApi.ApiGroup,
		ApiPath:   toAddApi.ApiPath,
		ApiMethod: toAddApi.ApiMethod,
	}
	if err := ServiceApi.AddApi(*addApi); err != nil {
		global.GqaLog.Error("添加API失败！", zap.Any("err", err))
		global.ErrorMessage("添加API失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加API成功！", c)
	}
}

func (a *ApiApi) DeleteApi(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceApi.DeleteApi(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除API失败！", zap.Any("err", err))
		global.ErrorMessage("删除API失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "删除API成功！")
		global.SuccessMessage("删除API成功！", c)
	}
}
