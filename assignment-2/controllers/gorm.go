package controllers

import (
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}
