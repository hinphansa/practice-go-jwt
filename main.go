package main

import (
	"fmt"
	"github/Hiinnn/practice-go/config"
	"github/Hiinnn/practice-go/models"
	"github/Hiinnn/practice-go/routes"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	dsn := config.GetConfig(config.BuildDBConfig())
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status :", err)
	}
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
