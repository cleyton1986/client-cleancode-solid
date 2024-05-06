package main

import (
	"log"

	"github.com/cleyton1986/client-cleancode-solid/internal/entities"
	"github.com/cleyton1986/client-cleancode-solid/internal/interfaces/controllers"
	"github.com/cleyton1986/client-cleancode-solid/internal/interfaces/repositories"
	"github.com/cleyton1986/client-cleancode-solid/internal/usecases"
	"github.com/cleyton1986/client-cleancode-solid/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
    // Inicializa o servidor web usando Gin
    r := gin.Default()

    // Inicializa a conexão com o banco de dados
    db, err := database.InitializeDatabase()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Garante que o banco de dados esteja atualizado com as migrações necessárias
    db.AutoMigrate(&entities.User{})

    // Cria instâncias dos repositórios
    userRepo := repositories.NewUserRepository(db)

    // Cria instâncias dos casos de uso
    userUseCase := usecases.NewUserUseCase(userRepo)

    // Cria instâncias dos controladores e configura as rotas
    userController := controllers.NewUserController(userUseCase)
    r.POST("/users", userController.CreateUser)
    r.GET("/users", userController.ListUsers)
    r.PUT("/users/:id", userController.UpdateUser)
    r.DELETE("/users/:id", userController.DeleteUser)

    // Inicia o servidor na porta 3333
    r.Run(":3333")
}
