package models

import "time"

type RoleHasPermission struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	RoleID       int
	Role         Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PermissionId int
	Permission   Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Created_at   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	Updated_at   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
