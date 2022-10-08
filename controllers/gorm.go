package controllers

import "gorm.io/gorm"

type InDB struct {
	db *gorm.DB
}

func NewInDB(db *gorm.DB) *InDB {
	return &InDB{db}
}
