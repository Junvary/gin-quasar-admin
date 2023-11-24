package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiMenu struct{}

func (a *ApiMenu) GetMenuList(c *gin.Context) {
	var toGetDataList model.RequestGetMenuList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceMenu.GetMenuList(toGetDataList); err != nil {
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

func (a *ApiMenu) EditMenu(c *gin.Context) {
	var toEditData model.SysMenu
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceMenu.EditMenu(c, toEditData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditSuccess"), c)
	}
}

func (a *ApiMenu) AddMenu(c *gin.Context) {
	var toAddData model.RequestAddMenu
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
	addData := &model.SysMenu{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddData.ParentCode,
		Name:                              toAddData.Name,
		Path:                              toAddData.Path,
		Component:                         toAddData.Component,
		Hidden:                            toAddData.Hidden,
		KeepAlive:                         toAddData.KeepAlive,
		Title:                             toAddData.Title,
		Icon:                              toAddData.Icon,
		IsLink:                            toAddData.IsLink,
	}
	if err := servicePrivate.ServiceMenu.AddMenu(*addData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiMenu) DeleteMenuById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceMenu.DeleteMenuById(c, toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiMenu) QueryMenuById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceMenu.QueryMenuById(toQueryId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}
