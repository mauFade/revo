package infra

import (
	"os"

	"github.com/joho/godotenv"
	postmodel "github.com/mauFade/revo/application/post/model"
	usermodel "github.com/mauFade/revo/application/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func ConnectToDatabase() {
	var err error

	connectionURI := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(connectionURI), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(
		usermodel.User{},
		postmodel.Post{},
	)
}
