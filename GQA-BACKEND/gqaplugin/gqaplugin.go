package gqaplugin

/*
    // ========================= Plugin development ===========================
	// Locally developed plugins can be ignored by .gitignore.
	// The developed plugins can be pushed to a separate repo.
    // ============================== Memo ====================================
    // To develop plugins, you need to implement the GqaPlugin interface.
    // And then fill it into the PluginList slice below.
	// Plugins are generally customized, and some features is not necessary.
    // For example, i18n.
	// ========================================================================
*/

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var PluginList = []GqaPlugin{
	example.PluginExample,
	achievement.PluginAchievement,
}

/*
	============================= GqaPlugin interface =================================
*/

type GqaPlugin interface {
	PluginCode() string                                     //Plugin code, used for routing packet name
	PluginName() string                                     //Plugin Name
	PluginVersion() string                                  //Plugin Version
	PluginMemo() string                                     //Plugin Memo
	PluginRouterPublic(publicGroup *gin.RouterGroup)        //Plugin Public Router
	PluginRouterPrivate(privateGroup *gin.RouterGroup)      //Plugin Private Router
	PluginMigrate() []interface{}                           //Plugin Migrations
	PluginData() []interface{ LoadData() (err error) }      //Plugin Default Data
	PluginCron() ([]gqaModel.SysCron, map[uuid.UUID]func()) // Plugin Cron
}

func RegisterPluginRouter(PublicGroup, PrivateGroup *gin.RouterGroup) {
	for _, p := range PluginList {
		PluginRouter(PublicGroup, PrivateGroup, p)
	}
}

func MigratePluginModel() []interface{} {
	var model []interface{}
	for _, p := range PluginList {
		model = append(model,
			p.PluginMigrate()...,
		)
	}
	return model
}

func LoadPluginData() []interface{ LoadData() (err error) } {
	var data []interface{ LoadData() (err error) }
	for _, p := range PluginList {
		data = append(data,
			p.PluginData()...,
		)
	}
	return data
}

func PluginRouter(publicGroup, privateGroup *gin.RouterGroup, Plugin ...GqaPlugin) {
	// Provides the plugin with a separate routing group named PluginCode()
	for i := range Plugin {
		PublicGroup := publicGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPublic(PublicGroup)
		PrivateGroup := privateGroup.Group(Plugin[i].PluginCode())
		Plugin[i].PluginRouterPrivate(PrivateGroup)
	}
}

func GetPluginCron() ([]gqaModel.SysCron, map[uuid.UUID]func()) {
	var pluginCronList []gqaModel.SysCron
	var pluginCronMap = make(map[uuid.UUID]func())
	for _, p := range PluginList {
		s, m := p.PluginCron()
		for _, sc := range s {
			pluginCronList = append(pluginCronList, sc)
			pluginCronMap[sc.UUID] = m[sc.UUID]
		}
	}
	return pluginCronList, pluginCronMap
}
