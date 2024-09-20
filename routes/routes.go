package routes

import (
	"go-todo-api-01/controllers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes inicializa todas as rotas da API
func InitializeRoutes(router *gin.Engine) {
	// Grupo de rotas de todo's
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("/", controllers.GetAllTodos)
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetTodoByID)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}
}
