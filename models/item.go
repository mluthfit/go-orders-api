package models

import (
	"time"
)

type ItemDefault struct {
	ItemCode    string `json:"itemCode" binding:"required,unique"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required,numeric"`
}

type Item struct {
	ItemID uint `json:"itemId" gorm:"primarykey"`
	ItemDefault
	OrderID   int       `json:"orderId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ItemUpdatePayload struct {
	LineItemID uint `json:"lineItemId" binding:"required"`
	ItemDefault
}
