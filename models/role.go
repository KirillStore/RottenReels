package models

type Role struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Desc  string
	Users []User `gorm:"foreignKey:RoleID"`
}
