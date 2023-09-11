package implement

import (
	"github.com/shafiyah/go-mux/models"
	repository "github.com/shafiyah/go-mux/repositories"
	services "github.com/shafiyah/go-mux/services"
)

type CustomerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func CreateCustomerService(customerRepo repository.CustomerRepository) services.CustomerService {
	return &CustomerServiceImpl{customerRepo}
}

// ReadAll implements services.CustomerService.
func (e *CustomerServiceImpl) ReadAll() (*[]models.Customer, error) {
	return e.customerRepo.FindAll()
}

// ReadById implements services.CustomerService.
func (e *CustomerServiceImpl) ReadById(id int64) (*models.Customer, error) {
	return e.customerRepo.ReadById(id)
}
