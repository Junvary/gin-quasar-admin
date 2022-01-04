package public_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/xk/service/public_service"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDownload(c *gin.Context)  {
	var getDownloadList model.RequestDownloadList
	if err := c.ShouldBindJSON(&getDownloadList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, news, total := public_service.GetDownload(getDownloadList); err!=nil{
		global.GqaLog.Error("获取资源列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取资源列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  news,
			Page:     getDownloadList.Page,
			PageSize: getDownloadList.PageSize,
			Total:    total,
		}, c)
	}
}
