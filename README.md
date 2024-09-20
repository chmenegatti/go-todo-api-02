## Construindo uma REST API com SOLID em Go - Parte 2: Banco de Dados, GORM e Persistência

Neste artigo, daremos continuidade à construção da nossa API RESTful, configurando o banco de dados SQLite, integrando o ORM GORM e implementando a persistência de dados, aplicando os princípios SOLID e a arquitetura TDD.

### Banco de Dados, GORM e Persistência
Persistir dados é uma parte essencial de qualquer aplicação. Hoje vamos configurar o banco de dados SQLite, integrar o GORM para manipulação de dados e implementar as operações CRUD para o modelo `Todo`.

#### Instalando o GORM

O GORM é um ORM (Object-Relational Mapping) para Go que simplifica a interação com o banco de dados. Para instalar o GORM, execute o seguinte comando:

```bash
go get gorm.io/gorm
```

#### Instalando o Driver SQLite

Para conectar ao banco de dados SQLite, precisamos instalar o driver correspondente. Execute o seguinte comando:

```bash
go get gorm.io/driver/sqlite
```

#### Passo 1: Definindo o Modelo

Na pasta `models`, vamos criar o modelo `Todo`. O modelo representa a estrutura dos dados que serão persistidos no banco.

**models/todo.go**:
```go
package models

import "gorm.io/gorm"

// Todo representa a estrutura da tabela no banco de dados
type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
```

O `gorm.Model` adiciona automaticamente campos como `ID`, `CreatedAt`, `UpdatedAt`, e `DeletedAt` ao modelo.

#### Passo 2: Configurando o Banco de Dados SQLite

Primeiro, vamos criar a conexão com o banco de dados SQLite. Adicione uma nova pasta chamada `database` e crie o arquivo `database.go`.

**database/database.go**:
```go
package database

import (
	"log"
	"go-todo-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conecta ao banco de dados e faz as migrações
func Connect() {
	// Conecta ao banco de dados SQLite
	database, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Auto-migração para criar a tabela no banco
	database.AutoMigrate(&models.Todo{})

	DB = database
}
```

Aqui, a função `Connect` inicializa a conexão com o SQLite e executa a migração automática do modelo `Todo` para criar a tabela no banco de dados.

#### Passo 3: Atualizando o Servidor para Conectar ao Banco de Dados

Agora, vamos garantir que o banco de dados seja inicializado quando o servidor for iniciado. Edite o arquivo `main.go` para incluir a conexão ao banco.

**main.go**:
```go
package main

import (
	"go-todo-api/database"
	"go-todo-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta ao banco de dados
	database.Connect()

	// Inicializa o router Gin
	router := gin.Default()

	// Configura as rotas
	routes.InitializeRoutes(router)

	// Sobe o servidor na porta 8080
	router.Run(":8080")
}
```

#### Passo 4: Implementando as Funções de Persistência no Controlador

Agora que o banco está configurado, vamos atualizar o controlador `Todo` para persistir os dados no SQLite.

**controllers/todo_controller.go**:
```go
package controllers

import (
	"go-todo-api/database"
	"go-todo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllTodos retorna todos os todos
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
```

Aqui, fizemos a integração com o GORM, utilizando-o para criar, ler, atualizar e deletar os "Todo's" diretamente no banco de dados.

#### Passo 5: Testando a Persistência de Dados

Agora, podemos testar nossa API. Certifique-se de que o servidor esteja rodando e faça as seguintes requisições usando `curl`:

1. **Criar um novo Todo**:
```bash
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"Estudar Go", "description":"Ler a documentação do Go", "completed":false}'
```

2. **Obter todos os Todo's**:
```bash
curl http://localhost:8080/todos
```

3. **Obter um Todo por ID**:
```bash
curl http://localhost:8080/todos/1
```

4. **Atualizar um Todo**:
```bash
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Estudar Go", "description":"Finalizar a série de artigos sobre Go", "completed":true}'
```

5. **Deletar um Todo**:
```bash
curl -X DELETE http://localhost:8080/todos/1
```

#### Conclusão

Neste artigo, configuramos o banco de dados SQLite, integramos o GORM para manipular a persistência de dados e implementamos as operações CRUD para o modelo `Todo`. No próximo artigo, continuaremos avançando na API com testes e refinamento da aplicação, sempre aplicando os princípios SOLID.

Fique atento para as próximas partes da série!