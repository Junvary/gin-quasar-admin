package gqacron

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/google/uuid"
	"time"
)

var uid = uuid.New()

var T1 = model.SysCron{
	UUID: uid,
	Id:   0,
	Name: "Gqa Test Task1",
	Spec: "@every 1m",
}

func CronTest() {
	t := time.Now().Format(time.DateTime)
	fmt.Println("Test cron from Gin-Quasar-Admin every 1m: ", t)
}
