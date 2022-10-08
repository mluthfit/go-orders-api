package models

import (
	"time"
)

type Order struct {
	OrderID uint `json:"orderId" gorm:"primaryKey"`
	OrderDefault
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
