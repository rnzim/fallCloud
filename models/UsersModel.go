package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"Name"`
	Password string `json:"Password"`
	Balance float64 `json:"Balance"`
    AcessLevel string `json:"AcessLevel"`
	
}
