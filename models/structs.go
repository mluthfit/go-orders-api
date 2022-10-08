package models

import "time"

type ItemDefault struct {
	ItemCode    string `json:"itemCode" binding:"required" gorm:"unique;not null"`
	Description string `json:"description" binding:"required" gorm:"not null"`
	Quantity    uint   `json:"quantity" binding:"required,numeric" gorm:"not null"`
}

type ItemUpdatePayload struct {
	LineItemID uint `json:"lineItemId" binding:"required"`
	ItemDefault
}

type OrderDefault struct {
	CustomerName string     `json:"customerName" binding:"required" gorm:"not null"`
	OrderedAt    *time.Time `json:"orderedAt" binding:"required" gorm:"not null"`
}

type OrderUpdatePayload struct {
	OrderDefault
	Items []ItemUpdatePayload `json:"items"`
}
