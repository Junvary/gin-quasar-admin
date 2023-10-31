package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"runtime"
)

type ApiDb struct{}

func (a *ApiDb) CheckDb(c *gin.Context) {
	goVersion := runtime.Version()
	ginVersion := gin.Version

	var pluginList []model.Plugin
	for _, v := range gqaplugin.PluginList {
		pluginList = append(pluginList, model.Plugin{
			PluginName:    v.PluginName(),
			PluginCode:    v.PluginCode(),
			PluginVersion: v.PluginVersion(),
			PluginMemo:    v.PluginMemo(),
		})
	}
	if global.GqaDb != nil {
		pluginLoginLayout := utils.GetConfigFrontend("pluginLoginLayout")
		model.ResponseSuccessMessageDataWithLog(gin.H{
			"need_init":           false,
			"go_version":          goVersion,
			"gin_version":         ginVersion,
			"plugin_list":         pluginList,
			"plugin_login_layout": pluginLoginLayout}, utils.GqaI18n(c, "DbNoNeedInit"), c)
		return
	} else {
		model.ResponseSuccessMessageDataWithLog(gin.H{
			"need_init":   true,
			"go_version":  goVersion,
			"gin_version": ginVersion,
			"plugin_list": pluginList}, utils.GqaI18n(c, "DbNeedInit"), c)
		return
	}
}

func (a *ApiDb) InitDb(c *gin.Context) {
	if global.GqaDb != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DbNoNeedInit"), c)
		return
	}

	var initDbInfo model.RequestDbInit
	if err := model.RequestShouldBindJSON(c, &initDbInfo); err != nil {
		return
	}
	if err := servicePublic.ServiceDb.InitDb(c, initDbInfo); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "CreateDbError")+", "+err.Error(), c)
		return
	}
	model.ResponseSuccessMessage(utils.GqaI18n(c, "InitSuccess"), c)
}
