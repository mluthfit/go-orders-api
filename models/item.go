package models

import (
	"time"
)

type ItemDefault struct {
	ItemCode    string `json:"itemCode" binding:"required" gorm:"unique;not null"`
	Description string `json:"description" binding:"required" gorm:"not null"`
	Quantity    uint   `json:"quantity" binding:"required,numeric" gorm:"not null"`
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
