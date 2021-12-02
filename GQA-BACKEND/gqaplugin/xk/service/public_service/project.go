package public_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetProject(getProjectList model.RequestProjectList) (err error, project []model.GqaPluginXkProject, total int64) {
	pageSize := getProjectList.PageSize
	offset := getProjectList.PageSize * (getProjectList.Page - 1)
	db := global.GqaDb.Model(&model.GqaPluginXkProject{})
	var projectList []model.GqaPluginXkProject
	//配置搜索
	if getProjectList.ProjectName != ""{
		db = db.Where("project_name like ?", "%" + getProjectList.ProjectName + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getProjectList.SortBy, getProjectList.Desc)).
		Preload("Leader").Find(&projectList).Error
	return err, projectList, total
}
