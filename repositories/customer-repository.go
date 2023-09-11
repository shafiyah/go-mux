package repository

import "github.com/shafiyah/go-mux/models"

type CustomerRepository interface {
	FindAll() (*[]models.Customer, error)
	ReadById(id int64) (*models.Customer, error)
}
