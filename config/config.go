package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Hulhay/goldfish/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabaseConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{LogLevel: logger.Info},
		),
	})
	if err != nil {
		panic("failed to create connection to database")
	}

	db.AutoMigrate(
		&model.User{},
		&model.Family{},
		&model.Member{},
		&model.Category{},
		&model.Wallet{},
		&model.Transaction{},
	)

	return db
}
