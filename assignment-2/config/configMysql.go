package config

import (
	"assignment-2/structs"

	"github.com/jinzhu/gorm"
)

// DBinit create conn to db
func DBInitMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3308)/orders_by?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed connect to mysql database")
	}
	db.AutoMigrate(structs.Order{}, structs.Item{})
	return db
}
