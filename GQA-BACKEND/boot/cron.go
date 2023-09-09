package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqacron"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/robfig/cron/v3"
)

func Cron() {
	c := cron.New()
	global.GqaCron = c
	global.GqaCron.Start()
	fmt.Println("Gin-Quasar-Admin Cron Start Succeeded!")
	global.GqaSLogger.Info("Gin-Quasar-Admin Cron Start Succeeded!")
	// Gin-Quasar-Admin cron
	gqacron.CronList = append(gqacron.CronList, gqacron.T1)
	gqacron.CronMap[gqacron.T1.UUID] = gqacron.CronTest
	// add more Gin-Quasar-Admin cron...

	// plugin cron
	pluginCronList, pluginCronMap := gqaplugin.GetPluginCron()
	for _, v := range pluginCronList {
		gqacron.CronList = append(gqacron.CronList, v)
		gqacron.CronMap[v.UUID] = pluginCronMap[v.UUID]
	}
}
