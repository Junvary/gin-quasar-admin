package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqacron"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiCron struct{}

func (a *ApiCron) StartCron(c *gin.Context) {
	startFlag := true
	var toQueryUuid model.RequestQueryByUUID
	if err := model.RequestShouldBindJSON(c, &toQueryUuid); err != nil {
		return
	}
	for i := 0; i < len(gqacron.CronList); i++ {
		if gqacron.CronList[i].UUID == toQueryUuid.UUID {
			if gqacron.CronList[i].Id != 0 {
				startFlag = false
				break
			}
			entryID, _ := global.GqaCron.AddFunc(gqacron.CronList[i].Spec, gqacron.CronMap[gqacron.CronList[i].UUID])
			gqacron.CronList[i].Id = entryID
			break
		}
	}
	if startFlag {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "CronStartSuccess"), c)
	} else {
		model.ResponseErrorMessage(utils.GqaI18n(c, "CronAlreadyStart"), c)
	}
}

func (a *ApiCron) StopCron(c *gin.Context) {
	stopFlag := true
	var toQueryUuid model.RequestQueryByUUID
	if err := model.RequestShouldBindJSON(c, &toQueryUuid); err != nil {
		return
	}
	for i := 0; i < len(gqacron.CronList); i++ {
		if gqacron.CronList[i].UUID == toQueryUuid.UUID {
			if gqacron.CronList[i].Id == 0 {
				stopFlag = false
				break
			}
			global.GqaCron.Remove(gqacron.CronList[i].Id)
			gqacron.CronList[i].Id = 0
			break
		}
	}
	if stopFlag {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "CronStopSuccess"), c)
	} else {
		model.ResponseErrorMessage(utils.GqaI18n(c, "CronAlreadyStop"), c)
	}
}

func (a *ApiCron) GetCronList(c *gin.Context) {
	model.ResponseSuccessData(gin.H{"records": gqacron.CronList}, c)
}
