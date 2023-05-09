package model

type SysTodo struct {
	GqaModelWithCreatedByAndUpdatedBy
	TodoDetail string `json:"todo_detail" gorm:"comment:内容;type:text;"`
	TodoStatus string `json:"todo_status" gorm:"comment:状态;default:yesNo_no;index;"`
}

type RequestGetTodoList struct {
	RequestPageAndSort
	TodoStatus string `json:"todo_status"`
}

type RequestAddTodo struct {
	TodoDetail string `json:"todo_detail"`
}
