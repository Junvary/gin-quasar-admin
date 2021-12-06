package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDict struct {}

func (a *ApiDict) GetDictList(c *gin.Context) {
	var requestDictList system.RequestDictList
	if err := c.ShouldBindJSON(&requestDictList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dictList, total, parentCode := ServiceDict.GetDictList(requestDictList); err != nil {
		global.GqaLog.Error("获取字典列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取字典列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePageWithParentId{
			Records:    dictList,
			Page:       requestDictList.Page,
			PageSize:   requestDictList.PageSize,
			Total:      total,
			ParentCode: parentCode,
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
	toEditDict.UpdatedBy = utils.GetUsername(c)
	if err := ServiceDict.EditDict(toEditDict); err != nil {
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
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDict.Status,
			Sort:      toAddDict.Sort,
			Remark:    toAddDict.Remark,
		},
		ParentCode: toAddDict.ParentCode,
		DictCode:   toAddDict.DictCode,
		DictLabel:  toAddDict.DictLabel,
	}
	if err := ServiceDict.AddDict(*addDict); err != nil {
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
	if err := ServiceDict.DeleteDict(toDeleteId.Id); err != nil {
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
	if err, dict := ServiceDict.QueryDictById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找字典失败！", zap.Any("err", err))
		global.ErrorMessage("查找字典失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dict}, "查找字典成功！", c)
	}
}
