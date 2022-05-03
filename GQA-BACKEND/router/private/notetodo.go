package private

import "github.com/gin-gonic/gin"

type RouterNoteTodo struct{}

func (r *RouterNoteTodo) InitRouterRouterNoteTodo(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("note-todo")
	{
		privateGroup.POST("get-note-todo-list", apiPrivate.ApiNoteTodo.GetNoteTodoList)
		privateGroup.POST("edit-note-todo", apiPrivate.ApiNoteTodo.EditNoteTodo)
		privateGroup.POST("add-note-todo", apiPrivate.ApiNoteTodo.AddNoteTodo)
		privateGroup.POST("delete-note-todo-by-id", apiPrivate.ApiNoteTodo.DeleteNoteTodoById)
		privateGroup.POST("query-note-todo-by-id", apiPrivate.ApiNoteTodo.QueryNoteTodoById)
	}
	return privateGroup
}
