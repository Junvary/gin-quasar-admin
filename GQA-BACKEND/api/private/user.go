package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiUser struct{}

func (a *ApiUser) GetUserList(c *gin.Context) {
	var toGetDataList model.RequestGetUserList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceUser.GetUserList(toGetDataList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiUser) EditUser(c *gin.Context) {
	var toEditData model.SysUser
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	if toEditData.Username == "admin" && toEditData.Status == "onOff_off" {
		model.ResponseErrorMessage(utils.GqaI18n("CantDisableAdmin"), c)
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.EditUser(toEditData); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiUser) AddUser(c *gin.Context) {
	var toAddData model.RequestAddUser
	if err := model.RequestShouldBindJSON(c, &toAddData); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddData.Status,
			Sort:      toAddData.Sort,
			Memo:      toAddData.Memo,
		},
	}
	addData := &model.SysUser{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		Avatar:                            toAddData.Avatar,
		Username:                          toAddData.Username,
		Nickname:                          toAddData.Nickname,
		RealName:                          toAddData.RealName,
		Gender:                            toAddData.Gender,
		Mobile:                            toAddData.Mobile,
		Email:                             toAddData.Email,
	}
	if err := servicePrivate.ServiceUser.AddUser(addData); err != nil {
		if err.Error() == "successWithNoDefaultPassword" {
			model.ResponseSuccessMessage(utils.GqaI18n("AddUserSuccessWithoutPwd"), c)
		} else {
			global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
			model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
		}
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiUser) DeleteUserById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	currentUsername := utils.GetUsername(c)
	err, currentUser := servicePrivate.ServiceUser.GetUserByUsername(currentUsername)
	if err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
		return
	}
	if currentUser.Id == toDeleteId.Id {
		global.GqaLogger.Error(utils.GetUsername(c) + utils.GqaI18n("CantDeleteYourself"))
		model.ResponseErrorMessage(utils.GqaI18n("CantDeleteYourself"), c)
		return
	}
	// 初始化时 admin 的 Id 为 1，这里就这样判断了，可以增加更多的逻辑。
	if toDeleteId.Id == 1 {
		global.GqaLogger.Error(utils.GetUsername(c) + utils.GqaI18n("CantDeleteSuperAdmin"))
		model.ResponseErrorMessage(utils.GqaI18n("CantDeleteSuperAdmin"), c)
		return
	}
	if err := servicePrivate.ServiceUser.DeleteUserById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GetUsername(c)+utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiUser) QueryUserById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceUser.QueryUserById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n("FindSuccess"), c)
	}
}

func (a *ApiUser) ResetPassword(c *gin.Context) {
	var toResetPasswordId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toResetPasswordId); err != nil {
		return
	}
	if err := servicePrivate.ServiceUser.ResetPassword(toResetPasswordId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("ResetPasswordFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("ResetPasswordFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("ResetPasswordSuccess"), c)
	}
}

func (a *ApiUser) GetUserMenu(c *gin.Context) {
	err, defaultPageList, menu, buttons := servicePrivate.ServiceUser.GetUserMenu(c)
	if err != nil {
		model.ResponseErrorMessage(utils.GqaI18n("GetUserMenuFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(
			gin.H{"records": menu, "buttons": buttons, "default_page_list": defaultPageList},
			utils.GqaI18n("GetUserMenuSuccess"), c,
		)
	}
}

func (a *ApiUser) ChangePassword(c *gin.Context) {
	var toChangePassword model.RequestChangePassword
	if err := model.RequestShouldBindJSON(c, &toChangePassword); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.ChangePassword(username, toChangePassword); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("ChangePasswordFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("ChangePasswordFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("ChangePasswordSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("ChangePasswordSuccess"), c)
	}
}

func (a *ApiUser) ChangeNickname(c *gin.Context) {
	var toChangeNickname model.RequestChangeNickname
	if err := model.RequestShouldBindJSON(c, &toChangeNickname); err != nil {
		return
	}
	if toChangeNickname.Nickname == "" {
		global.GqaLogger.Error(utils.GqaI18n("NullNickname"))
		model.ResponseErrorMessage(utils.GqaI18n("NullNickname"), c)
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceUser.ChangeNickname(username, toChangeNickname); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("ChangeNicknameFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("ChangeNicknameFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("ChangeNicknameSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("ChangeNicknameSuccess"), c)
	}
}
