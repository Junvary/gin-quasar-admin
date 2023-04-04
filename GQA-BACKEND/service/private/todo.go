package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type ServiceTodo struct{}

func (s *ServiceTodo) GetTodoList(requestTodoList model.RequestGetTodoList, username string) (err error, todoNote interface{}, total int64) {
	pageSize := requestTodoList.PageSize
	offset := requestTodoList.PageSize * (requestTodoList.Page - 1)
	db := global.GqaDb.Where("created_by = ?", username).Model(&model.SysTodo{})
	var todoNoteList []model.SysTodo
	// Search
	if requestTodoList.TodoStatus != "" {
		db = db.Where("todo_status = ?", requestTodoList.TodoStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestTodoList.SortBy, requestTodoList.Desc)).Find(&todoNoteList).Error
	return err, todoNoteList, total
}

func (s *ServiceTodo) EditTodo(toEditTodo model.SysTodo) (err error) {
	var sysTodo model.SysTodo
	if err = global.GqaDb.Where("id = ?", toEditTodo.Id).First(&sysTodo).Error; err != nil {
		return err
	}
	//err = global.GqaDb.Updates(&toEditTodo).Error
	err = global.GqaDb.Save(&toEditTodo).Error
	return err
}

func (s *ServiceTodo) AddTodo(toAddTodo model.SysTodo) (err error) {
	err = global.GqaDb.Create(&toAddTodo).Error
	return err
}

func (s *ServiceTodo) DeleteTodoById(id uint) (err error) {
	var sysTodo model.SysTodo
	if err = global.GqaDb.Where("id = ?", id).First(&sysTodo).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysTodo).Error
	return err
}

func (s *ServiceTodo) QueryTodoById(id uint) (err error, todoNoteInfo model.SysTodo) {
	var todoNote model.SysTodo
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&todoNote, "id = ?", id).Error
	return err, todoNote
}
