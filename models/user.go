package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique" json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	RoleID   uint     `gorm:"index" json:"role_id"`
	Role     Role     `gorm:"foreignKey:RoleID"`
	Movies   []Movie  `gorm:"foreignKey:UserID"`
	Ratings  []Rating `gorm:"foreignKey:UserID"`
	Reviews  []Review `gorm:"foreignKey:UserID"`
}
