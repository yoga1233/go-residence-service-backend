package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	loadEnv()

	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "go_residence")
	dbUser := GetEnv("DB_USER", "root")
	dbPass := GetEnv("DB_PASS", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.TenantOrder{}, &model.Report{}, &model.Tenant{}, &model.News{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected successfully")

}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv mengambil nilai variabel lingkungan dengan fallback jika tidak ada
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
