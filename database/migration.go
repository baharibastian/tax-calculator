package database

import (
	"github.com/tax-calculator/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Tax_object{})

	return
}