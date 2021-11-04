package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDict struct {
}

func (a *ApiDict) GetDictList(c *gin.Context) {
	var pageInfo system.RequestPageByParentId
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dictList, total, parentId := service.GroupServiceApp.ServiceSystem.GetDictList(pageInfo); err != nil {
		global.GqaLog.Error("获取字典列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取字典列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePageWithParentId{
			Records:  dictList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
			ParentId: parentId,
		}, c)
	}
}

func (a *ApiDict) EditDict(c *gin.Context) {
	var toEditDict system.SysDict
	if err := c.ShouldBindJSON(&toEditDict); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.EditDict(toEditDict); err != nil {
		global.GqaLog.Error("编辑字典失败!", zap.Any("err", err))
		global.ErrorMessage("编辑字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑字典成功！", c)
	}
}

func (a *ApiDict) AddDict(c *gin.Context) {
	var toAddDict system.RequestAddDict
	if err := c.ShouldBindJSON(&toAddDict); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addDict := &system.SysDict{
		GqaModel: global.GqaModel{
			Status: toAddDict.Status,
			Sort:   toAddDict.Sort,
			Remark: toAddDict.Remark,
		},
		ParentId: toAddDict.ParentId,
		Value:    toAddDict.Value,
		Label:    toAddDict.Label,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddDict(*addDict); err != nil {
		global.GqaLog.Error("添加字典失败！", zap.Any("err", err))
		global.ErrorMessage("添加字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加字典成功！", c)
	}
}

func (a *ApiDict) DeleteDict(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := service.GroupServiceApp.ServiceSystem.DeleteDict(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除字典失败！", zap.Any("err", err))
		global.ErrorMessage("删除字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除字典成功！", c)
	}
}

func (a *ApiDict) QueryDictById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dict := service.GroupServiceApp.ServiceSystem.QueryDictById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找字典失败！", zap.Any("err", err))
		global.ErrorMessage("查找字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dict}, "查找字典成功！", c)
	}
}

func (a *ApiDict) QueryDictByParentId(c *gin.Context) {
	var toQueryParentId system.RequestPageByParentId
	if err := c.ShouldBindJSON(&toQueryParentId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dict := service.GroupServiceApp.ServiceSystem.QueryDictByParentId(toQueryParentId.ParentId); err != nil {
		global.GqaLog.Error("查找字典失败！", zap.Any("err", err))
		global.ErrorMessage("查找字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dict}, "查找字典成功！", c)
	}
}
