package models

import (
	"time"
)

type Item struct {
	ItemID uint `json:"itemId" gorm:"primaryKey"`
	ItemDefault
	OrderID   int       `json:"orderId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
