package models

import (
	"time"
)

type OrderDefault struct {
	CustomerName string    `json:"customerName" binding:"required" gorm:"not null"`
	OrderedAt    time.Time `json:"orderedAt" binding:"required" gorm:"not null"`
}

type Order struct {
	OrderID uint `json:"orderId" gorm:"primarykey"`
	OrderDefault
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrderUpdatePayload struct {
	OrderDefault
	Items []ItemUpdatePayload `json:"items"`
}
