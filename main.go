package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shafiyah/go-mux/config"

	controller "github.com/shafiyah/go-mux/controllers"
	customerRepo "github.com/shafiyah/go-mux/repositories/implement"
	customerService "github.com/shafiyah/go-mux/services/implement"
)

func main() {
	config.ConnectDatabase()
	router := mux.NewRouter()

	customerRepo := customerRepo.CreateCustomerRepo(config.DB)
	customerService := customerService.CreateCustomerService(customerRepo)
	controller := controller.CreateCustomerController(customerService)

	router.HandleFunc("/customer", controller.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.GetCustomerById).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
