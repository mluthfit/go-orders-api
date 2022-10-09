package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderUpdatePayload struct {
	CustomerName string              `json:"customerName" binding:"required"`
	OrderedAt    *time.Time          `json:"orderedAt" binding:"required"`
	Items        []ItemUpdatePayload `json:"items"`
}

type Order struct {
	OrderID      uint       `json:"orderId" gorm:"primaryKey"`
	CustomerName string     `json:"customerName" binding:"required" gorm:"not null"`
	OrderedAt    *time.Time `json:"orderedAt" binding:"required" gorm:"not null"`
	Items        []Item     `json:"items"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

func (order *Order) GetAllOrdersAndItems(db *gorm.DB) (*[]Order, error) {
	var orders []Order

	var err = db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (order *Order) CreateOrderAndItems(db *gorm.DB) (*Order, error) {
	var tx = db.Begin()
	if err := tx.Omit("Items").Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for index := range order.Items {
		order.Items[index].OrderID = order.OrderID
		if err := tx.Create(&order.Items[index]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return order, nil
}

func (order *Order) UpdateOrderAndItems(db *gorm.DB, orderId uint, payloadOrder OrderUpdatePayload) (*Order, error) {
	if err := db.First(&order, orderId).Error; err != nil {
		return nil, err
	}

	var tx = db.Begin()
	if err := tx.Model(&order).
		Updates(Order{
			CustomerName: payloadOrder.CustomerName,
			OrderedAt:    payloadOrder.OrderedAt,
		}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var item Item
	for _, each := range payloadOrder.Items {
		if err := tx.Where(Item{
			OrderID: orderId,
			ItemID:  each.LineItemID,
		}).First(&item).
			Error; err != nil {
			tx.Rollback()
			return nil,
				fmt.Errorf(fmt.Sprintf("the item id %d was not found in the order", each.LineItemID))
		}

		if err := tx.Model(&item).Updates(each).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		order.Items = append(order.Items, item)
	}

	tx.Commit()
	return order, nil
}

func (order *Order) DeleteOrderAndItems(db *gorm.DB, orderId uint) (*Order, error) {
	if err := db.Preload("Items").First(&order, orderId).Error; err != nil {
		return nil, err
	}

	db.Select("Items").Delete(&order)
	return order, nil
}
