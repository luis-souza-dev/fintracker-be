package database

import (
	"fmt"
	"log"
	"os"

	"github.com/luis-souza-dev/fintracker-be/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// connectDb
func ConnectDb() {

	connectionStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	db.AutoMigrate(&models.Resident{}, &models.ExpensesCategories{}, &models.Expense{});

	DB = db
}