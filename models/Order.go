package models

import "time"

type Order struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID     uint64    `gorm:"not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID;constraint:onDelete:CASCADE" json:"user"`
	Items      string    `gorm:"type:json" json:"items"`
	Total      int       `gorm:"type:int;default:0" json:"total"`
	Status     int       `gorm:"type:tinyint(5);default:0" json:"status"`
	Comment    string    `gorm:"type:varchar(255)" json:"comment"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
