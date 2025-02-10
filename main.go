package main

import (
	"github.com/LandimTiago/gin-api-golang/models"
	"github.com/LandimTiago/gin-api-golang/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{
			Nome: "Tiago",
			CPF:  "123.123.123.12",
			RG:   "12.123.123.1",
		},
		{
			Nome: "Allana",
			CPF:  "213.213.213.21",
			RG:   "21.213.213.1",
		},
		{
			Nome: "Felipe",
			CPF:  "433.433.433.43",
			RG:   "43.433.433.1",
		},
		{
			Nome: "Ana",
			CPF:  "113.113.113.11",
			RG:   "11.113.113.1",
		},
	}

	routes.HandleRequest()
}
