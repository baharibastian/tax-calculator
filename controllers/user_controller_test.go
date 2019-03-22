package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/models"
	"github.com/ecojuntak/gorb/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	uu := make([]models.User, 1)
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Users").Return(uu)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.Resources)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
