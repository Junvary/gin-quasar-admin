package private

import "github.com/gin-gonic/gin"

type RouterTodo struct{}

func (r *RouterTodo) InitRouterRouterTodo(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
	privateGroup = privateGroup.Group("todo")
	{
		privateGroup.POST("get-todo-list", apiPrivate.ApiTodo.GetTodoList)
		privateGroup.POST("edit-todo", apiPrivate.ApiTodo.EditTodo)
		privateGroup.POST("add-todo", apiPrivate.ApiTodo.AddTodo)
		privateGroup.POST("delete-todo-by-id", apiPrivate.ApiTodo.DeleteTodoById)
		privateGroup.POST("query-todo-by-id", apiPrivate.ApiTodo.QueryTodoById)
	}
	return privateGroup
}
