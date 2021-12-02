package public_api

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
	"gin-quasar-admin/gqaplugin/xk/service/public_service"
	"gin-quasar-admin/model/system"
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
