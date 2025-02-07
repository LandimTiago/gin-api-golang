package controllers

import "github.com/gin-gonic/gin"

func RetornaAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"alunos": []string{"João", "Maria", "Pedro", "Ana"},
	})
}
