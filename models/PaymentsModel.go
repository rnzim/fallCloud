package models

import "gorm.io/gorm"


type Payments struct {
	gorm.Model
	Value float64 `json:"Value"`
	Status string `json:"Status`
	Method string `Json:"Method"`
	UserID  string  `json:"UserID"`

}