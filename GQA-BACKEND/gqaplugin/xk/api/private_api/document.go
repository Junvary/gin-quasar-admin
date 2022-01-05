package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/private_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDocumentList(c *gin.Context)  {
	var getDocumentList model.RequestDocumentList
	if err := c.ShouldBindJSON(&getDocumentList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, document, total := private_service.GetDocumentList(getDocumentList, utils.GetUsername(c)); err!=nil{
		global.GqaLog.Error("获取文档列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取文档列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  document,
			Page:     getDocumentList.Page,
			PageSize: getDocumentList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditDocument(c *gin.Context) {
	var toEditDocument model.GqaPluginXkDocument
	if err := c.ShouldBindJSON(&toEditDocument); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditDocument.UpdatedBy = utils.GetUsername(c)
	if err := private_service.EditDocument(toEditDocument, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("编辑文档失败！", zap.Any("err", err))
		global.ErrorMessage("编辑文档失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑文档成功！", c)
	}
}

func AddDocument(c *gin.Context) {
	var toAddDocument model.RequestAddDocument
	if err := c.ShouldBindJSON(&toAddDocument); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addDocument := &model.GqaPluginXkDocument{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDocument.Status,
			Sort:      toAddDocument.Sort,
			Remark:    toAddDocument.Remark,
		},
		Title: toAddDocument.Title,
		Content: toAddDocument.Content,
		Attachment: toAddDocument.Attachment,
	}
	if err := private_service.AddDocument(*addDocument, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加文档失败！", zap.Any("err", err))
		global.ErrorMessage("添加文档失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加文档成功！", c)
	}
}

func DeleteDocument(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.DeleteDocument(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除文档失败！", zap.Any("err", err))
		global.ErrorMessage("删除文档失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除文档成功！", c)
	}
}

func  QueryDocumentById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := private_service.QueryDocumentById(toQueryId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("查找文档失败！", zap.Any("err", err))
		global.ErrorMessage("查找文档失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找文档成功！", c)
	}
}
