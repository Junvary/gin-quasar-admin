package publicapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/service/publicservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

{{ range .PluginModel }}
{{ if .WithPublicList }}
func Get{{.ModelName}}List(c *gin.Context) {
	var toGetDataList model.RequestGet{{.ModelName}}List
	if err := gqaModel.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, records, total := publicservice.Get{{.ModelName}}List(toGetDataList); err != nil {
		gqaGlobal.GqaLogger.Error("获取{{.ModelName}}列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取{{.ModelName}}列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  records,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}
{{ end }}
{{ end }}