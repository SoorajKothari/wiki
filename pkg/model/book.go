package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}
