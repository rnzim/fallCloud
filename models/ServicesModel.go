package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	UserID string `json:"UserID"`
	Name string `json:"Name"`
	Status bool `json:"Status"`
	IdContainer string `json:"IdContainer"`

	
}