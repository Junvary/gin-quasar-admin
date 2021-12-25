package system

import "github.com/gin-gonic/gin"

type RouterTodoNote struct {}

func (r *RouterTodoNote) InitRouterTodoNote(Router *gin.RouterGroup) (R gin.IRoutes) {
	todoNoteGroup := Router.Group("todo-note")/*.Use(middleware.LogOperationHandler())*/
	{
		//获取待办便签列表
		todoNoteGroup.POST("todo-note-list", ApiTodoNote.GetTodoNoteList)
		//编辑待办便签信息
		todoNoteGroup.PUT("todo-note-edit", ApiTodoNote.EditTodoNote)
		//新增待办便签
		todoNoteGroup.POST("todo-note-add", ApiTodoNote.AddTodoNote)
		//删除待办便签
		todoNoteGroup.DELETE("todo-note-delete", ApiTodoNote.DeleteTodoNote)
		//根据ID查找待办便签
		todoNoteGroup.POST("todo-note-id", ApiTodoNote.QueryTodoNoteById)
	}
	return Router
}
