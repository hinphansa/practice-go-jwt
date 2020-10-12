package config

import (
	"fmt"

	"gorm.io/gorm"
)

// DB -> Use for calling a function of database
var DB *gorm.DB

// DBConfig -> Database configuration
type DBConfig struct {
	Host     string
	DBName   string
	Port     int
	User     string
	Password string
	SSLMode  string
	TimeZone string
}

// BuildDBConfig -> Create DB Config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		DBName:   "practice_go",
		Port:     5432,
		User:     "practice_go",
		Password: "181113Hk",
		SSLMode:  "disable",
		TimeZone: "Asia%2FBangkok",
	}
	return &dbConfig
}

// GetConfig -> Get Config String
func GetConfig(dbConfig *DBConfig) string {
	fmt.Println("PW:", dbConfig.Password)
	return fmt.Sprint(
		"host=", dbConfig.Host,
		" user=", dbConfig.User,
		" password=", dbConfig.Password,
		" dbname=", dbConfig.DBName,
		" port=", dbConfig.Port,
		" sslmode=", dbConfig.SSLMode,
		" TimeZone=", dbConfig.TimeZone,
	)
}
