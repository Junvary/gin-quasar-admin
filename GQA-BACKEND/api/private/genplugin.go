package private

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

type ApiGenPlugin struct{}

func (a *ApiGenPlugin) GetGenPluginList(c *gin.Context) {
	var toGetDataList model.RequestGetGenPluginList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceGenPlugin.GetGenPluginList(toGetDataList); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiGenPlugin) GenPlugin(c *gin.Context) {
	var sysGenPlugin model.SysGenPlugin
	if err := model.RequestShouldBindJSON(c, &sysGenPlugin); err != nil {
		return
	}
	if err := servicePrivate.ServiceGenPlugin.GenPlugin(&sysGenPlugin); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiGenPlugin) DeleteGenPluginById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceGenPlugin.DeleteGenPluginById(c, toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiGenPlugin) DownloadGenPluginById(c *gin.Context) {
	var toDownloadId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDownloadId); err != nil {
		return
	}
	err, pluginFile := servicePrivate.ServiceGenPlugin.DownloadGenPluginById(toDownloadId.Id)
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("message", url.QueryEscape(err.Error()))
	} else {
		lastIndex := strings.LastIndex(pluginFile.PluginFile, "/")
		fileName := pluginFile.PluginFile[lastIndex:]
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File(pluginFile.PluginFile)
	}
}
