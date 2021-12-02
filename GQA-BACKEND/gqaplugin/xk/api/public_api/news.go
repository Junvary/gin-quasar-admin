package public_api

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
	"gin-quasar-admin/gqaplugin/xk/service/public_service"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetNews(c *gin.Context)  {
	var getNewsList model.RequestNewsList
	if err := c.ShouldBindJSON(&getNewsList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, news, total := public_service.GetNews(getNewsList); err!=nil{
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
