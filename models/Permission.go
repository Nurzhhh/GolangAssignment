package models

import "time"

type Permission struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Action     string    `gorm:"type:varchar(255)" json:"action"`
	Model      string    `gorm:"type:varchar(255)" json:"model"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
