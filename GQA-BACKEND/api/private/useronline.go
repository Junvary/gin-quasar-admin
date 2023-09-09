package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiUserOnline struct{}

func (a *ApiUserOnline) GetUserOnlineList(c *gin.Context) {
	var toGetDataList model.RequestGetUserOnlineList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceUserOnline.GetUserOnlineList(toGetDataList); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "GetListFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUserOnline) KickUserOnlineByUsername(c *gin.Context) {
	var toKickUsername model.RequestQueryByUsername
	if err := model.RequestShouldBindJSON(c, &toKickUsername); err != nil {
		return
	}
	err := servicePrivate.ServiceUserOnline.KickOnlineUserByUsername(toKickUsername.Username)
	if err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "KickFailed"), "err", err)
		model.ResponseErrorMessage(utils.GqaI18n(c, "KickFailed")+err.Error(), c)
		return
	} else {
		global.GqaSLogger.Info(utils.GetUsername(c) + utils.GqaI18n(c, "KickSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n(c, "KickSuccess"), c)
	}
}
