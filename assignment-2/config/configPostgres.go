package config

import (
	"assignment-2/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5433"
	dbname   = "orders_by"
	db       *gorm.DB
	err      error
)

func DBInitPostgres() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("failed connect to postgres database")
	}
	db.Debug().AutoMigrate(structs.Order{}, structs.Item{})
	return db
}
