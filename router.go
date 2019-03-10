package main

import (
	"github.com/tax-calculator/controllers"
	"github.com/tax-calculator/middlewares"
	"github.com/tax-calculator/repositories"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func LoadRouter(db *gorm.DB) (r *mux.Router) {
	userRepo := repositories.NewUserRepo(db)
	taxObjectRepo := repositories.NewTaxObjectRepo(db)
	userController := controllers.NewUserController(userRepo)
	taxObjectController := controllers.NewTaxObjectController(taxObjectRepo)

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/users", userController.Resources).Methods("GET", "POST")
	v1.HandleFunc("/user/{id}", userController.Resources).Methods("GET")
	v1.HandleFunc("/tax_object", taxObjectController.Resources).Methods("POST")
	v1.HandleFunc("/tax_object/{user_id}", taxObjectController.Resources).Methods("GET")

	r.Use(middlewares.LoggerMidldlware)

	return
}