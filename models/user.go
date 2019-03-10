package models

import "time"

type User struct {
	ID        uint64     `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id" `
	Username  string	 `gorm:"column:username; type:varchar(50)" json:"username"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}