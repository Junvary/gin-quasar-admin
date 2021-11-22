package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/service/private_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetNews(c *gin.Context)  {
	err, news := private_service.GetNews()
	if err != nil {
		global.GqaLog.Error("鉴权：获取新闻列表失败！", zap.Any("err", err))
		global.ErrorMessage("鉴权：获取新闻列表失败！"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": news}, "鉴权：获取新闻列表成功！", c)
	}
}