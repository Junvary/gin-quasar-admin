package boot

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin"
	"github.com/robfig/cron/v3"
	"time"
)

var TaskList = map[string]func(){
	"@every 1m": CronTest,
}

func CronTest() {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Test cron from Gin-Quasar-Admin every 1m: ", t)
}

func Cron() *cron.Cron {
	c := cron.New()
	resultList := gqaplugin.RegisterPluginCron(TaskList)
	for k, v := range resultList {
		entryID, err := c.AddFunc(k, v)
		if err != nil {
			fmt.Println("cron start error:", entryID)
			return nil
		}
	}
	c.Start()
	fmt.Println("cron started! list: ", c.Entries())
	return c
}
