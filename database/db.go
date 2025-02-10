package database

import (
	"log"

	"github.com/LandimTiago/gin-api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados!")
	}

	DB.AutoMigrate(&models.Aluno{})
}
