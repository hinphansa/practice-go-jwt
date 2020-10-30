package main

import (
	"github/Hiinnn/practice-go/config"
	"github/Hiinnn/practice-go/models"
	"github/Hiinnn/practice-go/routes"
	"github/Hiinnn/practice-go/services"
	"io/ioutil"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var (
		err     error
		content []byte
	)

	// Config DB
	dsn := config.GetConfig(config.BuildDBConfig())
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	config.DB.AutoMigrate(&models.User{})

	// Set Secret Key
	content, err = ioutil.ReadFile("secret_key.txt")
	if err != nil {
		return
	}
	services.SetSecretKey(content)

	r := routes.SetupRouter()
	r.Run(":8080")
}
