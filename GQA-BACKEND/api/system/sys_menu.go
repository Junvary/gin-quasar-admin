package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiMenu struct {
}

func (a *ApiMenu) GetUserMenu(c *gin.Context) {
	err, menu := service.GroupServiceApp.ServiceSystem.GetUserMenu(c)
	if err != nil {
		global.ErrorMessage("获取用户菜单失败！", c)
	}
	global.SuccessData(system.ResponseMenu{
		Menu: menu,
	}, c)
}

func (a *ApiMenu) GetMenuList(c *gin.Context) {
	var pageInfo system.RequestPage
	_ = c.ShouldBindJSON(&pageInfo)
	if err, menuList, total := service.GroupServiceApp.ServiceSystem.GetMenuList(pageInfo); err != nil {
		global.GqaLog.Error("获取菜单列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取菜单列表失败！", c)
	} else {
		global.SuccessData(system.ResponsePage{
			List:     menuList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiMenu) EditMenu(c *gin.Context) {
	var toEditMenu system.SysMenu
	_ = c.ShouldBindJSON(&toEditMenu)
	if err := service.GroupServiceApp.ServiceSystem.EditMenu(toEditMenu); err != nil {
		global.GqaLog.Error("编辑菜单失败!", zap.Any("err", err))
		global.ErrorMessage("编辑菜单失败！", c)
	} else {
		global.SuccessMessage("编辑菜单成功！", c)
	}
}

func (a *ApiMenu) AddMenu(c *gin.Context) {
	var toAddMenu system.RequestAddMenu
	_ = c.ShouldBindJSON(&toAddMenu)
	addMenu := &system.SysMenu{
		ParentId:  toAddMenu.ParentId,
		Name:      toAddMenu.Name,
		Path:      toAddMenu.Path,
		Component: toAddMenu.Component,
		Hidden:    toAddMenu.Hidden,
		KeepAlive: toAddMenu.KeepAlive,
		Title:     toAddMenu.Title,
		Icon:      toAddMenu.Icon,
		IsLink:    toAddMenu.IsLink,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddMenu(*addMenu); err != nil {
		global.GqaLog.Error("添加菜单失败！", zap.Any("err", err))
		global.ErrorMessage("添加菜单失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加菜单成功！", c)
	}
}

func (a *ApiMenu) DeleteMenu(c *gin.Context) {
	var toDeleteId system.RequestDelete
	_ = c.ShouldBindJSON(&toDeleteId)
	if err := service.GroupServiceApp.ServiceSystem.DeleteMenu(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除菜单失败！", zap.Any("err", err))
		global.ErrorMessage("删除菜单失败！", c)
	} else {
		global.SuccessMessage("删除菜单成功！", c)
	}
}

func (a *ApiMenu) QueryMenuById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	_ = c.ShouldBindJSON(&toQueryId)
	if err, menu := service.GroupServiceApp.ServiceSystem.QueryMenuById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找菜单失败！", zap.Any("err", err))
		global.ErrorMessage("查找菜单失败！", c)
	} else {
		global.SuccessMessageData(gin.H{"info": menu}, "查找菜单成功！", c)
	}
}
