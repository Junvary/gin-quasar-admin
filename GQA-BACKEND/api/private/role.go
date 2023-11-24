package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiRole struct{}

func (a *ApiRole) GetRoleList(c *gin.Context) {
	var toGetDataList model.RequestGetRoleList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceRole.GetRoleList(toGetDataList); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
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
	if err := servicePrivate.ServiceRole.EditRole(c, toEditData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditSuccess"), c)
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
	if err := servicePrivate.ServiceRole.AddRole(c, *addData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiRole) DeleteRoleById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.DeleteRoleById(toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiRole) QueryRoleById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceRole.QueryRoleById(toQueryId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}

func (a *ApiRole) GetRoleMenuList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, menuList := servicePrivate.ServiceRole.GetRoleMenuList(&roleCode); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetRoleMenuListFailed")+err.Error(), c)
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
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditRoleMenuFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditRoleMenuSuccess"), c)
	}
}

func (a *ApiRole) GetRoleApiList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, apiList := servicePrivate.ServiceRole.GetRoleApiList(&roleCode); err != nil {
		model.ResponseSuccessMessageWithLog(utils.GqaI18n(c, "GetRoleApiListFailed")+err.Error(), c)
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
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditRoleApiFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditRoleApiSuccess"), c)
	}
}

func (a *ApiRole) QueryUserByRole(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceRole.QueryUserByRole(&roleCode); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindRoleUserFailed")+err.Error(), c)
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
		model.ResponseErrorMessage(utils.GqaI18n(c, "CantRemoveAdminFromAdmin"), c)
		return
	}
	if err := servicePrivate.ServiceRole.RemoveRoleUser(&toDeleteRoleUser); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "RemoveRoleUserFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "RemoveRoleUserSuccess"), c)
	}
}

func (a *ApiRole) AddRoleUser(c *gin.Context) {
	var toAddDataUser model.RequestRoleUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddDataUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.AddRoleUser(&toAddDataUser); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddRoleUserFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddRoleUserSuccess"), c)
	}
}

func (a *ApiRole) EditRoleDeptDataPermission(c *gin.Context) {
	var toEditDataDeptDataPermission model.RequestRoleDeptDataPermission
	if err := model.RequestShouldBindJSON(c, &toEditDataDeptDataPermission); err != nil {
		return
	}
	if err := servicePrivate.ServiceRole.EditRoleDeptDataPermission(&toEditDataDeptDataPermission); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditRoleDeptDataPermissionFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditRoleDeptDataPermissionSuccess"), c)
	}
}

func (a *ApiRole) GetRoleButtonList(c *gin.Context) {
	var roleCode model.RequestRoleCode
	if err := model.RequestShouldBindJSON(c, &roleCode); err != nil {
		return
	}
	if err, buttonList := servicePrivate.ServiceRole.GetRoleButtonList(&roleCode); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetRoleButtonFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": buttonList}, c)
	}
}
