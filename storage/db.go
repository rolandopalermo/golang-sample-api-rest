package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sample-api-rest/config"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	var err error
	dsn := config.GetSqlConnectionString()

	log.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
