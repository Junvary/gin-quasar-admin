package {{.PluginCode}}

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/router/private_router"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/router/public_router"
	"github.com/gin-gonic/gin"
)

var Plugin{{.PluginCode}} = new({{.PluginCode}})

type {{.PluginCode}} struct{}

func (f *{{.PluginCode}}) PluginCode() string {
	return "plugin-{{.PluginCode}}"
}

func (f *{{.PluginCode}}) PluginName() string {
	return "{{.PluginName}}"
}

func (f *{{.PluginCode}}) PluginVersion() string {
	return "v0.0.1"
}

func (f *{{.PluginCode}}) PluginRemark() string {
	return "这是{{.PluginName}}插件"
}

func (f *{{.PluginCode}}) PluginRouterPublic(publicGroup *gin.RouterGroup) {
    public_router.InitPublicRouter(publicGroup)
}

func (f *{{.PluginCode}}) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	private_router.InitPrivateRouter(privateGroup)
}

func (f *{{.PluginCode}}) PluginMigrate() []interface{} {
	return nil
}

func (f *{{.PluginCode}}) PluginData() []interface{ LoadData() (err error) } {
	return nil
}

func init() {
}
