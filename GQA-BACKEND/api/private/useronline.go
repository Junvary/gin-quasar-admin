package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUserOnline struct{}

func (a *ApiUserOnline) GetUserOnlineList(c *gin.Context) {
	var requestUserOnlineList model.RequestGetUserOnlineList
	if err := model.RequestShouldBindJSON(c, &requestUserOnlineList); err != nil {
		return
	}
	if err, userList, total := servicePrivate.ServiceUserOnline.GetUserOnlineList(requestUserOnlineList); err != nil {
		global.GqaLogger.Error("获取在线用户列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取在线用户列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  userList,
			Page:     requestUserOnlineList.Page,
			PageSize: requestUserOnlineList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUserOnline) KickUserOnlineByUsername(c *gin.Context) {
	var toDeleteUsername model.RequestQueryByUsername
	if err := model.RequestShouldBindJSON(c, &toDeleteUsername); err != nil {
		return
	}
	currentUsername := utils.GetUsername(c)
	err := servicePrivate.ServiceUserOnline.KickOnlineUserByUsername(currentUsername)
	if err != nil {
		global.GqaLogger.Error("踢出用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("踢出用户失败，"+err.Error(), c)
		return
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "踢出用户成功！")
		model.ResponseSuccessMessage("踢出用户成功！", c)
	}
}
