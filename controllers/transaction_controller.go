package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID int `json:"id"`
	Amount float64 `json:"amount"`
	Type string `json:"type"` // "income" or " expense"
}

// Simulação de banco de memória
var transactions = []Transaction{
	{ID: 1, Amount: 100.0, Type: "income"},
	{ID: 2, Amount: 50.0, Type: "expense"},
}

// GET /transactions 
func ListTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, transactions)
}

// POST /transactions
func CreateTransaction(c *gin.Context) {
	var t Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = len(transactions) + 1 // Simula um ID único
	transactions = append(transactions, t)
	c.JSON(http.StatusCreated, t)
}

// PUT /transactions/:id
func UpdateTransaction(c *gin.Context) {
	// Implementar a lógica de atualização
	c.JSON(http.StatusOK, gin.H{"message": "Atualizar transação"})
}

// DELETE /transactions/:id
func DeleteTransaction(c *gin.Context) {
	// Implemente a lógica de deletar a transação
	c.JSON(http.StatusOK, gin.H{"message": "Deletar transação"})
}
