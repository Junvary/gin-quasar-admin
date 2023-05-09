package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		global.GqaLogger.Info(utils.GqaI18n("DbNoNeedInit"))
		model.ResponseSuccessMessageData(gin.H{
			"need_init":           false,
			"go_version":          goVersion,
			"gin_version":         ginVersion,
			"plugin_list":         pluginList,
			"plugin_login_layout": pluginLoginLayout}, utils.GqaI18n("DbNoNeedInit"), c)
		return
	} else {
		global.GqaLogger.Info(utils.GqaI18n("DbNeedInit"))
		model.ResponseSuccessMessageData(gin.H{
			"need_init":   true,
			"go_version":  goVersion,
			"gin_version": ginVersion,
			"plugin_list": pluginList}, utils.GqaI18n("DbNeedInit"), c)
		return
	}
}

func (a *ApiDb) InitDb(c *gin.Context) {
	if global.GqaDb != nil {
		global.GqaLogger.Error(utils.GqaI18n("DbNoNeedInit"))
		model.ResponseErrorMessage(utils.GqaI18n("DbNoNeedInit"), c)
		return
	}

	var initDbInfo model.RequestDbInit
	if err := model.RequestShouldBindJSON(c, &initDbInfo); err != nil {
		return
	}
	if err := servicePublic.ServiceDb.InitDb(initDbInfo); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("CreateDbError"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("CreateDbError")+", "+err.Error(), c)
		return
	}
	model.ResponseSuccessMessage(utils.GqaI18n("InitSuccess"), c)
}
