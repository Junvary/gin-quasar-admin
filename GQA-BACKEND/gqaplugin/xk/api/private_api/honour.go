package private_api

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
	"gin-quasar-admin/gqaplugin/xk/service/private_service"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHonourList(c *gin.Context)  {
	var getHonourList model.RequestHonourList
	if err := c.ShouldBindJSON(&getHonourList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, honour, total := private_service.GetHonourList(getHonourList); err!=nil{
		global.GqaLog.Error("获取荣誉认证列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取荣誉认证列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  honour,
			Page:     getHonourList.Page,
			PageSize: getHonourList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditHonour(c *gin.Context) {
	var toEditHonour model.GqaPluginXkHonour
	if err := c.ShouldBindJSON(&toEditHonour); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditHonour.UpdatedBy = utils.GetUsername(c)
	if err := private_service.EditHonour(toEditHonour); err != nil {
		global.GqaLog.Error("编辑荣誉认证失败！", zap.Any("err", err))
		global.ErrorMessage("编辑荣誉认证失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑荣誉认证成功！", c)
	}
}

func AddHonour(c *gin.Context) {
	var toAddHonour model.RequestAddHonour
	if err := c.ShouldBindJSON(&toAddHonour); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addHonour := &model.GqaPluginXkHonour{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddHonour.Status,
			Sort:      toAddHonour.Sort,
			Remark:    toAddHonour.Remark,
		},
		Title: toAddHonour.Title,
		Attachment: toAddHonour.Attachment,
	}
	if err := private_service.AddHonour(*addHonour); err != nil {
		global.GqaLog.Error("添加荣誉认证失败！", zap.Any("err", err))
		global.ErrorMessage("添加荣誉认证失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加荣誉认证成功！", c)
	}
}

func DeleteHonour(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.DeleteHonour(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除荣誉认证失败！", zap.Any("err", err))
		global.ErrorMessage("删除荣誉认证失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除荣誉认证成功！", c)
	}
}

func  QueryHonourById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := private_service.QueryHonourById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找荣誉认证失败！", zap.Any("err", err))
		global.ErrorMessage("查找荣誉认证失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找荣誉认证成功！", c)
	}
}
