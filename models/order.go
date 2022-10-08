package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID uint `json:"orderId" gorm:"primaryKey"`
	OrderDefault
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (order *Order) GetAllOrdersAndItems(db *gorm.DB) (*[]Order, error) {
	var orders []Order

	var err = db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (order *Order) CreateOrderAndItems(db *gorm.DB) {

}

func (order *Order) UpdateOrderAndItems(db *gorm.DB) {

}

func (order *Order) DeleteOrderAndItems(db *gorm.DB, orderId uint) (*Order, error) {
	if err := db.First(&order, orderId).Error; err != nil {
		return nil, err
	}

	db.Select("Items").Delete(order)
	return order, nil
}
