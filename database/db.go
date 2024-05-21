package database

import (
	"log"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	//HOST := os.Getenv("HOST")
	//USER := os.Getenv("USER")
	//PASSWORD := os.Getenv("PASSWORD")
	//DBNAME := os.Getenv("DBNAME")
	//PORT := os.Getenv("PORT")

	HOST := "localhost"
	USER := "root"
	PASSWORD := "root"
	DBNAME := "root"
	PORT := "5432"

	stringDeConexao := "host=" + HOST +
		" user=" + USER +
		" password=" + PASSWORD +
		" dbname=" + DBNAME +
		" port=" + PORT +
		" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
