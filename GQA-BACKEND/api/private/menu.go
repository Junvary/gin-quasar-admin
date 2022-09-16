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
		global.GqaLogger.Error("获取菜单列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取菜单列表失败，"+err.Error(), c)
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
		global.GqaLogger.Error("编辑菜单失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑菜单失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑菜单成功！")
		model.ResponseSuccessMessage("编辑菜单成功！", c)
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
		global.GqaLogger.Error("添加菜单失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加菜单失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加菜单成功！", c)
	}
}

func (a *ApiMenu) DeleteMenuById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceMenu.DeleteMenuById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除菜单失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除菜单失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除菜单成功！")
		model.ResponseSuccessMessage("删除菜单成功！", c)
	}
}

func (a *ApiMenu) QueryMenuById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, menu := servicePrivate.ServiceMenu.QueryMenuById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找菜单失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找菜单失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": menu}, "查找菜单成功！", c)
	}
}
