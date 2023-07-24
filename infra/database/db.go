package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	strConnection := "user=postgres dbname=rest_api password=123456789 host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnection))
	if err != nil {
		log.Panic(err.Error())
	}
}
