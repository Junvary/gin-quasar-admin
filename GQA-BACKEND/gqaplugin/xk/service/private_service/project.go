package private_service

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/model"
)

func GetProjectList(getProjectList model.RequestProjectList) (err error, project []model.GqaPluginXkProject, total int64) {
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

func  EditProject(toEditProject model.GqaPluginXkProject) (err error) {
	var project model.GqaPluginXkProject
	if err = global.GqaDb.Where("id = ?", toEditProject.Id).First(&project).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditProject).Error
	return err
}

func AddProject(toAddProject model.GqaPluginXkProject) (err error) {
	err = global.GqaDb.Create(&toAddProject).Error
	return err
}

func DeleteProject(id uint) (err error) {
	var project model.GqaPluginXkProject
	if err = global.GqaDb.Where("id = ?", id).First(&project).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&project).Error
	return err
}

func QueryProjectById(id uint) (err error, projectInfo model.GqaPluginXkProject) {
	var project model.GqaPluginXkProject
	err = global.GqaDb.Preload("Leader").Preload("CreatedByUser").Preload("UpdatedByUser").First(&project, "id = ?", id).Error
	return err, project
}