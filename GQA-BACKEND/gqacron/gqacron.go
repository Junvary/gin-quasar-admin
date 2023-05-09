package gqacron

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/google/uuid"
)

var CronList []model.SysCron
var CronMap = make(map[uuid.UUID]func())
