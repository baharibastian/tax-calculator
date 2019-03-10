package controllers

import (
	"net/http"
	"strconv"
	"time"
	"github.com/tax-calculator/models"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/tax-calculator/repositories"
)

type TaxObjectController struct {
	repo repositories.TaxObjectRepository
}

type TaxObjectWithInfo struct {
	ID        uint64     `json:"id" `
	User_id	  int		 `json:"user_id"`
	Name      string	 `json:"name"`
	Tax_code  string     `json:"tax_code"`
	Price	  int        `json:"price"`
	Refundable bool		 `json:"refundable"`
	Tax       float32	 `json:"tax"`
	Amount    float32    `json:"amount"`
}

func NewTaxObjectController(repo repositories.TaxObjectRepository) TaxObjectController {
	return TaxObjectController{repo}
}

func (c *TaxObjectController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) > 0 {
			c.TaxObjects(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (c *TaxObjectController) TaxObjects(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logrus.Infoln(params)
	user_id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	tax_objects := c.repo.TaxObjects(user_id)

	var uu []TaxObjectWithInfo
	var total float32 = 0.0
	var tax_subtotal float32 = 0.0
	var price_subtotal = 0

	for _, tax_object := range tax_objects {
		var refundable bool = false
		var tax float32 = 0.0

		if tax_object.Tax_code == 1 {
			refundable = true
			tax = float32(float64(tax_object.Price) * 0.1)
		}

		if tax_object.Tax_code == 2 {
			tax = float32((float64(tax_object.Price) * 0.02) + 10)
		}

		if tax_object.Tax_code == 3 {
			if tax_object.Price >= 100 {
				tax = float32((float64(tax_object.Price) - 100) * 0.01)
			}
		}
		row := TaxObjectWithInfo{
			ID:			tax_object.ID,
			User_id: 	tax_object.User_id,
			Name: 		tax_object.Name,
			Tax_code: 	taxCodeString(tax_object.Tax_code),
			Price: 		tax_object.Price,
			Refundable: refundable,
			Tax: 		tax,
			Amount: 	float32(tax_object.Price) + tax,
		}

		price_subtotal += tax_object.Price
		tax_subtotal += tax
		total += row.Amount

		uu = append(uu, row)
	}

	response := map[string]interface{}{
		"tax_object_items": uu,
		"price_subtotal": price_subtotal,
		"tax_subtotal": tax_subtotal,
		"total": total,
	}

	respondWithJSON(w, http.StatusOK, response)
}

func (c *TaxObjectController) Create(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	tax_code, err := strconv.Atoi(r.FormValue("tax_code"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user_id, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := models.Tax_object{
		Name: r.FormValue("name"),
		User_id: user_id,
		Tax_code: tax_code,
		Price: price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	user := c.repo.Store(u)

	respondWithJSON(w, http.StatusCreated, user)
}

func taxCodeString(tax_code int) (string) {
	tax_name := [3]string{"Food & Beverage", "Tobacco", "Entertainment"}

	return tax_name[tax_code-1 ]
}