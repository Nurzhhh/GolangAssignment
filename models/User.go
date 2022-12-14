package models

import "time"

type User struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Token      string `gorm:"-" json:"token,omitempty"`
	RoleID     int
	Role       *Role     `gorm:"foreignkey:RoleID;constraint:OnDelete:SET NULL"`
	Orders     []Order   `json:"orders,omitempty"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
