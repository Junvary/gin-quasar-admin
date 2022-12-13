package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceNoteTodo struct{}

func (s *ServiceNoteTodo) GetNoteTodoList(requestNoteTodoList model.RequestGetNoteTodoList, username string) (err error, todoNote interface{}, total int64) {
	pageSize := requestNoteTodoList.PageSize
	offset := requestNoteTodoList.PageSize * (requestNoteTodoList.Page - 1)
	db := global.GqaDb.Where("created_by = ?", username).Model(&model.SysNoteTodo{})
	var todoNoteList []model.SysNoteTodo
	// Search
	if requestNoteTodoList.TodoStatus != "" {
		db = db.Where("todo_status = ?", requestNoteTodoList.TodoStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestNoteTodoList.SortBy, requestNoteTodoList.Desc)).Find(&todoNoteList).Error
	return err, todoNoteList, total
}

func (s *ServiceNoteTodo) EditNoteTodo(toEditNoteTodo model.SysNoteTodo) (err error) {
	var sysNoteTodo model.SysNoteTodo
	if err = global.GqaDb.Where("id = ?", toEditNoteTodo.Id).First(&sysNoteTodo).Error; err != nil {
		return err
	}
	//err = global.GqaDb.Updates(&toEditNoteTodo).Error
	err = global.GqaDb.Save(&toEditNoteTodo).Error
	return err
}

func (s *ServiceNoteTodo) AddNoteTodo(toAddNoteTodo model.SysNoteTodo) (err error) {
	err = global.GqaDb.Create(&toAddNoteTodo).Error
	return err
}

func (s *ServiceNoteTodo) DeleteNoteTodoById(id uint) (err error) {
	var sysNoteTodo model.SysNoteTodo
	if err = global.GqaDb.Where("id = ?", id).First(&sysNoteTodo).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysNoteTodo).Error
	return err
}

func (s *ServiceNoteTodo) QueryNoteTodoById(id uint) (err error, todoNoteInfo model.SysNoteTodo) {
	var todoNote model.SysNoteTodo
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&todoNote, "id = ?", id).Error
	return err, todoNote
}
