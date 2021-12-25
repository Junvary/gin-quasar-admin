package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type ServiceTodoNote struct{}

func (s *ServiceTodoNote) GetTodoNoteList(requestTodoNoteList system.RequestTodoNoteList, username string) (err error, role interface{}, total int64) {
	pageSize := requestTodoNoteList.PageSize
	offset := requestTodoNoteList.PageSize * (requestTodoNoteList.Page - 1)
	db := global.GqaDb.Where("created_by = ?", username).Model(&system.SysTodoNote{})
	var todoNoteList []system.SysTodoNote
	//配置搜索
	if requestTodoNoteList.TodoStatus != "" {
		db = db.Where("todo_status = ?", requestTodoNoteList.TodoStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestTodoNoteList.SortBy, requestTodoNoteList.Desc)).Find(&todoNoteList).Error
	return err, todoNoteList, total
}

func (s *ServiceTodoNote) EditTodoNote(toEditTodoNote system.SysTodoNote) (err error) {
	var sysTodoNote system.SysTodoNote
	if err = global.GqaDb.Where("id = ?", toEditTodoNote.Id).First(&sysTodoNote).Error; err != nil {
		return err
	}
	err = global.GqaDb.Updates(&toEditTodoNote).Error
	return err
}

func (s *ServiceTodoNote) AddTodoNote(toAddTodoNote system.SysTodoNote) (err error) {
	err = global.GqaDb.Create(&toAddTodoNote).Error
	return err
}

func (s *ServiceTodoNote) DeleteTodoNote(id uint) (err error) {
	var sysTodoNote system.SysTodoNote
	if err = global.GqaDb.Where("id = ?", id).First(&sysTodoNote).Error; err != nil {
		return err
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysTodoNote).Error
	return err
}

func (s *ServiceTodoNote) QueryTodoNoteById(id uint) (err error, todoNoteInfo system.SysTodoNote) {
	var todoNote system.SysTodoNote
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&todoNote, "id = ?", id).Error
	return err, todoNote
}
