package models

import (
	"time"
)

type OrderDefault struct {
	CustomerName string    `json:"customerName" binding:"required"`
	OrderedAt    time.Time `json:"orderedAt" binding:"required"`
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
