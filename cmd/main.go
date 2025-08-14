package main

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/thiago-tertuliano/Financa-API/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"log"
	"github.com/thiago-tertuliano/Financa-API/config"

)

// @title Financa API
// @version 1.0
// @descriptions API de Controle Financeiro em GO
// @host localhost:8080
// @BasePath /api/v1

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	// Rotas públicas
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// Rotas de autenticação
	}

	// Rotas Protegidas (exemplo, depois adicionar middleware de autenticação)
	// auth := v1.Group("/")
	// auth.Use(AuthMiddleware())
	// {
	//     auth.GET("/transactions", ListTransactions)
	//     auth.POST("/transactions", CreateTransaction)
	//	   // etc.
	//}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Servidor rodando na porta %s", cfg.Port)
    r.Run(":" + cfg.Port)
}