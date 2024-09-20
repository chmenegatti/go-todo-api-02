package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Retorna todos os todos"})
}

// CreateTodo cria um novo todo
func CreateTodo(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Todo criado"})
}

// GetTodoByID retorna um todo pelo ID
func GetTodoByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Retorna um todo pelo ID"})
}

// UpdateTodo atualiza um todo
func UpdateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Todo atualizado"})
}

// DeleteTodo deleta um todo
func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Todo deletado"})
}
