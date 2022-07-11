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
	var get{{.ModelName}}List model.RequestGet{{.ModelName}}List
	if err := gqaModel.RequestShouldBindJSON(c, &get{{.ModelName}}List); err != nil {
		return
	}
	if err, records, total := publicservice.Get{{.ModelName}}List(get{{.ModelName}}List); err != nil {
		gqaGlobal.GqaLogger.Error("获取{{.ModelName}}列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取{{.ModelName}}列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  records,
			Page:     get{{.ModelName}}List.Page,
			PageSize: get{{.ModelName}}List.PageSize,
			Total:    total,
		}, c)
	}
}
{{ end }}
{{ end }}