package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tax-calculator/models"

	"github.com/gorilla/mux"

	"github.com/tax-calculator/repositories"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) UserController {
	return UserController{repo}
}

func (c *UserController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Users(w, r)
		} else {
			c.User(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (c *UserController) Users(w http.ResponseWriter, r *http.Request) {
	users := c.repo.Users()

	var uu []models.User

	for _, user := range users {
		uu = append(uu, user)
	}

	respondWithJSON(w, http.StatusOK, uu)
}

func (c *UserController) User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := c.repo.User(id)

	respondWithJSON(w, http.StatusOK, u)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	u := models.User{
		Username:  r.FormValue("username"),
		CreatedAt: now,
		UpdatedAt: now,
	}

	user := c.repo.Create(u)

	respondWithJSON(w, http.StatusCreated, user)
}
