package repositories

import (
	"github.com/tax-calculator/models"
	"github.com/jinzhu/gorm"
)

type taxObjectRepo struct {
	DB *gorm.DB
}

type TaxObjectRepository interface {
	TaxObjects(user_id int) []models.Tax_object
	Store(u models.Tax_object) models.Tax_object
}

func NewTaxObjectRepo(DB *gorm.DB) TaxObjectRepository {
	return &taxObjectRepo{DB}
}

func (r *taxObjectRepo) TaxObjects(user_id int) []models.Tax_object {
	var uu []models.Tax_object

	// r.DB.Find(&uu)
	r.DB.Where("user_id = ?", user_id).Find(&uu)
	return uu
}

func (r *taxObjectRepo) Store(u models.Tax_object) models.Tax_object {
	r.DB.Create(&u)

	return u
}