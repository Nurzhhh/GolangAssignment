package models

import "time"

type Product struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int       `gorm:"type:int;default:0" json:"price"`
	Status      bool      `gorm:"type:bool;default:0" json:"status"`
	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
