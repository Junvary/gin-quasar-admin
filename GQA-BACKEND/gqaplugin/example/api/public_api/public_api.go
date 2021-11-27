package public_api

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/example/service/public_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetNews(c *gin.Context)  {
	err, news := public_service.GetNews()
	if err != nil {
		global.GqaLog.Error("公开：获取新闻列表失败！", zap.Any("err", err))
		global.ErrorMessage("公开：获取新闻列表失败！"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": news}, "公开：获取新闻列表成功！", c)
	}
}
