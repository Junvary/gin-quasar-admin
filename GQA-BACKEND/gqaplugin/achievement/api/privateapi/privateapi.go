package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/achievement/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	var getCategoryList model.RequestGetCategoryList
	if err := gqaModel.RequestShouldBindJSON(c, &getCategoryList); err != nil {
		return
	}
	if err, records, total := privateservice.GetCategoryList(getCategoryList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("获取Category列表失败！", "err", err)
		gqaModel.ResponseErrorMessage("获取Category列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  records,
			Page:     getCategoryList.Page,
			PageSize: getCategoryList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditCategory(c *gin.Context) {
	var toEditCategory model.PluginAchievementCategory
	if err := gqaModel.RequestShouldBindJSON(c, &toEditCategory); err != nil {
		return
	}
	toEditCategory.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditCategory(toEditCategory, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("编辑Category失败！", "err", err)
		gqaModel.ResponseErrorMessage("编辑Category失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("编辑Category成功！", c)
	}
}

func AddCategory(c *gin.Context) {
	var toAddCategory model.RequestAddCategory
	if err := gqaModel.RequestShouldBindJSON(c, &toAddCategory); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddCategory.Status,
			Sort:      toAddCategory.Sort,
			Memo:      toAddCategory.Memo,
		},
	}
	addCategory := &model.PluginAchievementCategory{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
	}
	if err := privateservice.AddCategory(c, *addCategory, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("添加Category失败！", "err", err)
		gqaModel.ResponseErrorMessage("添加Category失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("添加Category成功！", c)
	}
}

func DeleteCategoryById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteCategoryById(c, toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("删除Category失败！", "err", err)
		gqaModel.ResponseErrorMessage("删除Category失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessage("删除Category成功！", c)
	}
}

func QueryCategoryById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, record := privateservice.QueryCategoryById(c, toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("查找Category失败！", "err", err)
		gqaModel.ResponseErrorMessage("查找Category失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": record}, "查找Category成功！", c)
	}
}

func GetObtainList(c *gin.Context) {
	var getObtainList model.RequestGetObtainList
	if err := gqaModel.RequestShouldBindJSON(c, &getObtainList); err != nil {
		return
	}
	if err, records, total := privateservice.GetObtainList(getObtainList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaSLogger.Error("获取Obtain列表失败！", "err", err)
		gqaModel.ResponseErrorMessage("获取Obtain列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  records,
			Page:     getObtainList.Page,
			PageSize: getObtainList.PageSize,
			Total:    total,
		}, c)
	}
}

func ObtainFind(c *gin.Context) {
	var obtainFind model.ObtainFind
	if err := gqaModel.RequestShouldBindJSON(c, &obtainFind); err != nil {
		return
	}
	find := privateservice.ObtainFind(obtainFind)
	if find {
		gqaModel.ResponseSuccessMessage("已获得该成就！", c)
	} else {
		toObtain := model.PluginAchievementObtain{
			CategoryCode: obtainFind.CategoryCode,
			CreatedBy:    obtainFind.Username,
		}
		ObtainAchievement(c, toObtain)
	}
}

func ObtainAchievement(c *gin.Context, toObtain model.PluginAchievementObtain) {
	if err := privateservice.ObtainAchievement(toObtain); err != nil {
		gqaGlobal.GqaSLogger.Error("记录成就失败！", "err", err)
		gqaModel.ResponseErrorMessage("记录成就失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gin.H{"get": true}, c)
	}
}
