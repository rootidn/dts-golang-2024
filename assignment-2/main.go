package main

import (
	"assignment-2/config"
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInitPostgres()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.POST("/order", inDB.CreateOrder)
	router.PUT("/order", inDB.UpdateOrder)
	router.DELETE("order/:id", inDB.DeleteOrder)
	router.Run(":3000")
}
