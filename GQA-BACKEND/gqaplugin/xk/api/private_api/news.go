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

func GetNewsList(c *gin.Context)  {
	var getNewsList model.RequestNewsList
	if err := c.ShouldBindJSON(&getNewsList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, news, total := private_service.GetNewsList(getNewsList, utils.GetUsername(c)); err!=nil{
		global.GqaLog.Error("获取最新要闻列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取最新要闻列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  news,
			Page:     getNewsList.Page,
			PageSize: getNewsList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditNews(c *gin.Context) {
	var toEditNews model.GqaPluginXkNews
	if err := c.ShouldBindJSON(&toEditNews); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditNews.UpdatedBy = utils.GetUsername(c)
	if err := private_service.EditNews(toEditNews, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("编辑最新要闻失败！", zap.Any("err", err))
		global.ErrorMessage("编辑最新要闻失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑最新要闻成功！", c)
	}
}

func AddNews(c *gin.Context) {
	var toAddNews model.RequestAddNews
	if err := c.ShouldBindJSON(&toAddNews); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addNews := &model.GqaPluginXkNews{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddNews.Status,
			Sort:      toAddNews.Sort,
			Remark:    toAddNews.Remark,
		},
		Title: toAddNews.Title,
		Content: toAddNews.Content,
		Attachment: toAddNews.Attachment,
	}
	if err := private_service.AddNews(*addNews, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加最新要闻失败！", zap.Any("err", err))
		global.ErrorMessage("添加最新要闻失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加最新要闻成功！", c)
	}
}

func DeleteNews(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.DeleteNews(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除最新要闻失败！", zap.Any("err", err))
		global.ErrorMessage("删除最新要闻失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除最新要闻成功！", c)
	}
}

func  QueryNewsById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := private_service.QueryNewsById(toQueryId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("查找最新要闻失败！", zap.Any("err", err))
		global.ErrorMessage("查找最新要闻失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找最新要闻成功！", c)
	}
}
