package controller

import (
	"desafio-itau-golang/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var transaction []models.Transaction
var nextId = 1

func CreateTransaction(c *gin.Context) {
	var newTransaction models.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(422, gin.H{"Não foi possivel cadastrar o usuario!": err.Error()})
		return
	}

	if newTransaction.Value <= 0 {
		err := errors.New("o valor não pode ser menor ou igual a zero")
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	if newTransaction.Date.After(time.Now()) {
		err := errors.New("A data não pode ser no futuro")
		c.JSON(422, gin.H{"error": err.Error()})
		return
	} else if newTransaction.Date.IsZero() {
		newTransaction.Date = time.Now()
	}

	newTransaction.ID = nextId
	nextId++
	transaction = append(transaction, newTransaction)

	c.JSON(201, newTransaction)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	idInt := 0
	if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for index, trans := range transaction {
		if trans.ID == idInt {
			transaction = append(transaction[:index], transaction[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Todas as informações foram apagadas com sucesso!",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
}

func GetStatistics() {

}
