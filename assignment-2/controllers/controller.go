package controllers

import (
	"assignment-2/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// create new order to db
func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	ordered_at_str := c.PostForm("orderedAt")
	customer_name := c.PostForm("customerName")
	items_json := c.PostForm("items")

	ordered_at, err := time.Parse(time.RFC3339, ordered_at_str)
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "orderedAt data format is wrong!",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	var items []structs.Item
	err = json.Unmarshal([]byte(items_json), &items)
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "item list format is wrong",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	order.Ordered_At = ordered_at
	order.Customer_Name = customer_name
	order.Items = items

	err = idb.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "gorm error when creating data",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"status": "success",
		"data":   order,
	}
	c.JSON(http.StatusCreated, result)
}

// get all orders
func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	err := idb.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "gorm error when finding all data",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	if len(orders) <= 0 {
		result = gin.H{
			"status": "success",
			"data":   nil,
		}
	} else {
		result = gin.H{
			"status": "success",
			"data":   orders,
		}
	}
	c.JSON(http.StatusOK, result)
}

// update order by {id}
func (idb *InDB) UpdateOrder(c *gin.Context) {
	id := c.Query("id")
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)

	ordered_at_str := c.PostForm("orderedAt")
	customer_name := c.PostForm("customerName")
	items_json := c.PostForm("items")

	ordered_at, _ := time.Parse(time.RFC3339, ordered_at_str)

	var items []structs.Item
	err := json.Unmarshal([]byte(items_json), &items)
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "item list format is wrong",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	newOrder.Ordered_At = ordered_at
	newOrder.Customer_Name = customer_name
	newOrder.Items = items
	idInt, _ := strconv.ParseUint(id, 10, 32)
	newOrder.ID = uint(idInt)

	err = idb.DB.Preload("Items").First(&order, idInt).Error
	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "item list format is wrong",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	if len(newOrder.Items) > len(order.Items) {
		result = gin.H{
			"status":  "failed",
			"message": "items length is not equal to update",
			"desc":    fmt.Sprintf("selected data have %d items while updating %d", len(order.Items), len(newOrder.Items)),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	for i, item := range order.Items {
		if i >= len(newOrder.Items) {
			break
		}
		newOrder.Items[i].ID = item.ID
	}

	err = idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&newOrder).Error
	// use this to save orders when item nums is more than the original (create new item)
	// err = idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&newOrder).Error

	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "gorm error when updating data",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"status": "success",
		"data":   newOrder,
	}
	c.JSON(http.StatusOK, result)
}

// delete order with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&order, id).Error

	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "data not found!",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	idb.DB.Clauses(clause.Returning{}).Where("order_id = ?", order.ID).Delete(&structs.Item{})
	err = idb.DB.Delete(&order).Error

	if err != nil {
		result = gin.H{
			"status":  "failed",
			"message": "gorm error when deleting data",
			"desc":    err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"status": "success",
			"data":   nil,
		}
	}
	c.JSON(http.StatusOK, result)
}
