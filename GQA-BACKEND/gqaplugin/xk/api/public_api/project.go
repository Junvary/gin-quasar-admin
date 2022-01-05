package public_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/public_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetProject(c *gin.Context)  {
	var getProjectList model.RequestProjectList
	if err := c.ShouldBindJSON(&getProjectList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, project, total := public_service.GetProject(getProjectList); err!=nil{
		global.GqaLog.Error("获取项目列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取项目列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  project,
			Page:     getProjectList.Page,
			PageSize: getProjectList.PageSize,
			Total:    total,
		}, c)
	}
}
