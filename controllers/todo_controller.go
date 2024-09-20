package controllers

import (
	"go-todo-api-02/database"
	"go-todo-api-02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// CreateTodo cria um novo todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// GetTodoByID retorna um todo pelo ID
func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo não encontrado"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo atualiza um todo
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo deleta um todo
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo não encontrado"})
		return
	}
	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deletado com sucesso"})
}
