package model

type SysNoteTodo struct {
	GqaModelWithCreatedByAndUpdatedBy
	TodoDetail string `json:"todo_detail" gorm:"comment:内容;type:text;"`
	TodoStatus string `json:"todo_status" gorm:"comment:状态;default:yesNo_no;index;"`
}

type RequestGetNoteTodoList struct {
	RequestPageAndSort
	TodoStatus string `json:"todo_status"`
}

type RequestAddNoteTodo struct {
	TodoDetail string `json:"todo_detail"`
}
