package models

import "time"

type Tax_object struct {
	ID        uint64     `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id" `
	User_id	  int		 `gorm:"column:user_id; size:5" json:"user_id"`
	Name      string	 `gorm:"column:name; type:varchar(100)" json:"name"`
	Tax_code  int        `gorm:"column:tax_code; size:20" json:"tax_code"`
	Price	  int        `gorm:"column:price; size:20" json:"price"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}