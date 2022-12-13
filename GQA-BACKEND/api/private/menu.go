package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiMenu struct{}

func (a *ApiMenu) GetMenuList(c *gin.Context) {
	var requestMenuList model.RequestGetMenuList
	if err := model.RequestShouldBindJSON(c, &requestMenuList); err != nil {
		return
	}
	if err, menuList, total := servicePrivate.ServiceMenu.GetMenuList(requestMenuList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  menuList,
			Page:     requestMenuList.Page,
			PageSize: requestMenuList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiMenu) EditMenu(c *gin.Context) {
	var toEditMenu model.SysMenu
	if err := model.RequestShouldBindJSON(c, &toEditMenu); err != nil {
		return
	}
	toEditMenu.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceMenu.EditMenu(toEditMenu); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiMenu) AddMenu(c *gin.Context) {
	var toAddMenu model.RequestAddMenu
	if err := model.RequestShouldBindJSON(c, &toAddMenu); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddMenu.Status,
			Sort:      toAddMenu.Sort,
			Memo:      toAddMenu.Memo,
		},
	}
	addMenu := &model.SysMenu{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddMenu.ParentCode,
		Name:                              toAddMenu.Name,
		Path:                              toAddMenu.Path,
		Component:                         toAddMenu.Component,
		Hidden:                            toAddMenu.Hidden,
		KeepAlive:                         toAddMenu.KeepAlive,
		Title:                             toAddMenu.Title,
		Icon:                              toAddMenu.Icon,
		IsLink:                            toAddMenu.IsLink,
	}
	if err := servicePrivate.ServiceMenu.AddMenu(*addMenu); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiMenu) DeleteMenuById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceMenu.DeleteMenuById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiMenu) QueryMenuById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, menu := servicePrivate.ServiceMenu.QueryMenuById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": menu}, utils.GqaI18n("FindSuccess"), c)
	}
}
