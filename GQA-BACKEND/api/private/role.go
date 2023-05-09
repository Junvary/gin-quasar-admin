package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiRole struct{}

func (a *ApiRole) GetRoleList(c *gin.Context) {
	var toGetDataList model.RequestGetRoleList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceRole.GetRoleList(toGetDataList); err != nil {
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

func (a *ApiRole) EditRole(c *gin.Context) {
	var toEditData model.SysRole
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceRole.EditRole(toEditData); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiRole) AddRole(c *gin.Context) {
	var toAddData model.RequestAddRole
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
	addData := &model.SysRole{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		RoleCode:                          toAddData.RoleCode,
		RoleName:                          toAddData.RoleName,
	}
	if err := servicePrivate.ServiceRole.AddRole(*addData); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiRole) DeleteRoleById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.DeleteRoleById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiRole) QueryRoleById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceRole.QueryRoleById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n("FindSuccess"), c)
	}
}

func (a *ApiRole) GetRoleMenuList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, menuList := servicePrivate.ServiceRole.GetRoleMenuList(&roleCode); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetRoleMenuListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetRoleMenuListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": menuList}, c)
	}
}

func (a *ApiRole) EditRoleMenu(c *gin.Context) {
	var roleMenu model.RequestRoleMenuEdit
	if err := model.RequestShouldBindJSON(c, &roleMenu); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleMenu(&roleMenu); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditRoleMenuFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditRoleMenuFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditRoleMenuSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditRoleMenuSuccess"), c)
	}
}

func (a *ApiRole) GetRoleApiList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, apiList := servicePrivate.ServiceRole.GetRoleApiList(&roleCode); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetRoleApiListFailed"), zap.Any("err", err))
		model.ResponseSuccessMessage(utils.GqaI18n("GetRoleApiListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": apiList}, c)
	}
}

func (a *ApiRole) EditRoleApi(c *gin.Context) {
	var roleApi model.RequestEditRoleApi
	if err := model.RequestShouldBindJSON(c, &roleApi); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleApi(&roleApi); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditRoleApiFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditRoleApiFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditRoleApiSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditRoleApiSuccess"), c)
	}
}

func (a *ApiRole) QueryUserByRole(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceRole.QueryUserByRole(&roleCode); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindRoleUserFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindRoleUserFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiRole) RemoveRoleUser(c *gin.Context) {
	var toDeleteRoleUser model.RequestRoleUser
	if err := model.RequestShouldBindJSON(c, &toDeleteRoleUser); err != nil {
		return
	}
	if toDeleteRoleUser.Username == "admin" && toDeleteRoleUser.RoleCode == "super-admin" {
		model.ResponseErrorMessage(utils.GqaI18n("CantRemoveAdminFromAdmin"), c)
		return
	}
	if err := servicePrivate.ServiceRole.RemoveRoleUser(&toDeleteRoleUser); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("RemoveRoleUserFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("RemoveRoleUserFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("RemoveRoleUserSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("RemoveRoleUserSuccess"), c)
	}
}

func (a *ApiRole) AddRoleUser(c *gin.Context) {
	var toAddDataUser model.RequestRoleUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddDataUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.AddRoleUser(&toAddDataUser); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddRoleUserFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddRoleUserFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddRoleUserSuccess"), c)
	}
}

func (a *ApiRole) EditRoleDeptDataPermission(c *gin.Context) {
	var toEditDataDeptDataPermission model.RequestRoleDeptDataPermission
	if err := model.RequestShouldBindJSON(c, &toEditDataDeptDataPermission); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleDeptDataPermission(&toEditDataDeptDataPermission); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditRoleDeptDataPermissionFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditRoleDeptDataPermissionFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditRoleDeptDataPermissionSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditRoleDeptDataPermissionSuccess"), c)
	}
}

func (a *ApiRole) GetRoleButtonList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, buttonList := servicePrivate.ServiceRole.GetRoleButtonList(&roleCode); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetRoleButtonFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetRoleButtonFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": buttonList}, c)
	}
}
