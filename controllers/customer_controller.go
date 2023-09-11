package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/shafiyah/go-mux/helper"
	"github.com/shafiyah/go-mux/models"
	"github.com/shafiyah/go-mux/services"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

type CustomerController struct {
	customerService services.CustomerService
}

func CreateCustomerController(customerService services.CustomerService) CustomerController {
	customerController := CustomerController{customerService}
	return customerController

}

func (e *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers *[]models.Customer
	customers, err := e.customerService.ReadAll()
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseJson(w, http.StatusOK, customers)
}

func (e *CustomerController) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var customer *models.Customer
	customer, err = e.customerService.ReadById(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, " customer not fount")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, customer)
}
