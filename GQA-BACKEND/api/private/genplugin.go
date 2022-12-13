package private

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

type ApiGenPlugin struct{}

func (a *ApiGenPlugin) GenPlugin(c *gin.Context) {
	var sysGenPlugin model.SysGenPlugin
	if err := model.RequestShouldBindJSON(c, &sysGenPlugin); err != nil {
		return
	}
	err := servicePrivate.ServiceGenPlugin.GenPlugin(&sysGenPlugin)
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("message", url.QueryEscape(err.Error()))
		_ = os.Remove("./gqa-gen-plugin.zip")

	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip"))
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.Header("abc", "abc")
		c.File("./gqa-gen-plugin.zip")
		_ = os.Remove("./gqa-gen-plugin.zip")
	}
}
