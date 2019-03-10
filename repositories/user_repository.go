package repositories

import (
	"github.com/tax-calculator/models"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	User(id int) models.User
	Users() []models.User
	Create(u models.User) models.User
}

func NewUserRepo(DB *gorm.DB) UserRepository {
	return &userRepo{DB}
}

func (r *userRepo) User(id int) models.User {
	var u models.User

	r.DB.Find(&u, id)

	return u
}

func (r *userRepo) Users() []models.User {
	var uu []models.User

	r.DB.Find(&uu)

	return uu
}

func (r *userRepo) Create(u models.User) models.User {
	r.DB.Create(&u)

	return u
}