package main

import (
	"go-todo-api-02/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o router Gin
	router := gin.Default()

	// Configura as rotas
	routes.InitializeRoutes(router)

	// Sobe o servidor na porta 8080
	router.Run(":8080")
}
