package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json: "username" binding: "required"`
	Password string `json: "password" binding: "required"`
}

type LoginInput struct {
	Username string `json: "username" binding: "required"`
	Password string `json: "password" binding: "required"`
}

// POST /register
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Aqui voce pode adicionar a lógica para registrar o usuário no banco de dados
	c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucessso!"})
}

// POST /login
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Aqui você faria a lógica de autenticação e geração de token JWT
	c.JSON(http.StatusOK, gin.H{"token": "jwt_token_placeholder"})
}