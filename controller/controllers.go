package controllers

import (
	"github.com/LandimTiago/gin-api-golang/models"
	"github.com/gin-gonic/gin"
)

func RetornaAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func RetornaSaudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"mensagem": "Olá " + nome + ", tudo bem?",
	})
}
