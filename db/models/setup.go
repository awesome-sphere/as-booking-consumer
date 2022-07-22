package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/awesome-sphere/as-booking-consumer/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var THEATER_AMOUNT int

func InitDatabase() {
	dbUser := utils.GetenvOr("POSTGRES_USER", "pkinwza")
	dbPassword := utils.GetenvOr("POSTGRES_PASSWORD", "securepassword")
	dbHost := utils.GetenvOr("POSTGRES_HOST", "127.0.0.1")
	dbPort := utils.GetenvOr("POSTGRES_PORT", "5432")
	dbName := utils.GetenvOr("POSTGRES_DB", "as-cinema")
	theater_amount := utils.GetenvOr("THEATER_AMOUNT", "5")
	theater_number, err := strconv.Atoi(theater_amount)
	if err != nil {
		log.Println("Fail to convert type of THEATER_AMOUNT")
	}
	THEATER_AMOUNT = theater_number

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println("Database Connection:" + err.Error())
	}
	DB = db
}
