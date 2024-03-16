package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_Name string    `json:"customerName"`
	Ordered_At    time.Time `json:"orderedAt"`
	Items         []Item    `gorm:"foreignKey:Order_ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}

type Item struct {
	gorm.Model
	Item_Code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    uint
}

// kunci jawaban https://go.dev/play/p/zEjdZOM_mjl
