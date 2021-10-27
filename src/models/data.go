package models

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Uuid string `gorm:"not null; unique"`
	Data string
}

/*

 */
