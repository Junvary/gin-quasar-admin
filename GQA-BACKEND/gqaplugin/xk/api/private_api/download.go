package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/private_service"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDownloadList(c *gin.Context)  {
	var getDownloadList model.RequestDownloadList
	if err := c.ShouldBindJSON(&getDownloadList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, download, total := private_service.GetDownloadList(getDownloadList, utils.GetUsername(c)); err!=nil{
		global.GqaLog.Error("获取资源列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取资源列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  download,
			Page:     getDownloadList.Page,
			PageSize: getDownloadList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditDownload(c *gin.Context) {
	var toEditDownload model.GqaPluginXkDownload
	if err := c.ShouldBindJSON(&toEditDownload); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditDownload.UpdatedBy = utils.GetUsername(c)
	if err := private_service.EditDownload(toEditDownload, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("编辑资源失败！", zap.Any("err", err))
		global.ErrorMessage("编辑资源失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑资源成功！", c)
	}
}

func AddDownload(c *gin.Context) {
	var toAddDownload model.RequestAddDownload
	if err := c.ShouldBindJSON(&toAddDownload); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addDownload := &model.GqaPluginXkDownload{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDownload.Status,
			Sort:      toAddDownload.Sort,
			Remark:    toAddDownload.Remark,
		},
		Title: toAddDownload.Title,
		Content: toAddDownload.Content,
		Attachment: toAddDownload.Attachment,
	}
	if err := private_service.AddDownload(*addDownload, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加资源失败！", zap.Any("err", err))
		global.ErrorMessage("添加资源失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加资源成功！", c)
	}
}

func DeleteDownload(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.DeleteDownload(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除资源失败！", zap.Any("err", err))
		global.ErrorMessage("删除资源失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除资源成功！", c)
	}
}

func  QueryDownloadById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := private_service.QueryDownloadById(toQueryId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("查找资源失败！", zap.Any("err", err))
		global.ErrorMessage("查找资源失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找资源成功！", c)
	}
}
