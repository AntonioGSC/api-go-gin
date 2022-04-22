package database

import (
	"log"
	"os"

	"github.com/AntonioGSC/api-go-gin/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")

	stringDeConexao := "host=localhost user=" + postgresUser + " password=" + postgresPassword + " dbname=" + postgresDB + " port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o DB")
	}

	DB.AutoMigrate(&models.Aluno{})
}
