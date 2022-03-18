package system

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"os"
)

type ApiGenPlugin struct{}

func (a *ApiGenPlugin) GenPlugin(c *gin.Context) {
	var sysGenPlugin system.SysGenPlugin
	if err := c.ShouldBindJSON(&sysGenPlugin); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	err := ServiceGenPlugin.GenPlugin(&sysGenPlugin)
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		_ = os.Remove("./gqa-gen-plugin.zip")

	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip"))
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./gqa-gen-plugin.zip")
		_ = os.Remove("./gqa-gen-plugin.zip")
	}
}
