package services

import "github.com/shafiyah/go-mux/models"

type CustomerService interface {
	ReadAll() (*[]models.Customer, error)
	ReadById(id int64) (*models.Customer, error)
}
