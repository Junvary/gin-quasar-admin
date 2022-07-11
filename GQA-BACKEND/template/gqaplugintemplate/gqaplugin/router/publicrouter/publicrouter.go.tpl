package publicrouter

import (
    "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/api/publicapi"
    "github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	{
	    {{ range .PluginModel }}
	    {{ if .WithPublicList }}
	    //获取{{.ModelName}}列表
        publicGroup.POST("get-{{.ModelName}}-list", publicapi.Get{{.ModelName}}List)
        {{ end }}
	    {{ end }}

	}
}

