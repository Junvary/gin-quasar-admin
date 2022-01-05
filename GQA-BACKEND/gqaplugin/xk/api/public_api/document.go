package public_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/public_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDocument(c *gin.Context)  {
	var getDocumentList model.RequestDocumentList
	if err := c.ShouldBindJSON(&getDocumentList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, news, total := public_service.GetDocument(getDocumentList); err!=nil{
		global.GqaLog.Error("获取文档列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取文档列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  news,
			Page:     getDocumentList.Page,
			PageSize: getDocumentList.PageSize,
			Total:    total,
		}, c)
	}
}
