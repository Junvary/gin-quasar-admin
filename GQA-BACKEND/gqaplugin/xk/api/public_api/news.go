package public_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/public_service"
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
