package main

import (
	"fmt"
	"github/Hiinnn/practice-go/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	dsn := config.GetConfig(config.BuildDBConfig())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	r.Run(":8080")
}
