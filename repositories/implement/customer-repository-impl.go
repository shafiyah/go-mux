package implement

import (
	"fmt"

	"github.com/shafiyah/go-mux/config"
	"github.com/shafiyah/go-mux/models"
	repository "github.com/shafiyah/go-mux/repositories"
	"gorm.io/gorm"
)

type CustomerRepoImpl struct {
	DB *gorm.DB
}

func CreateCustomerRepo(DB *gorm.DB) repository.CustomerRepository {
	return &CustomerRepoImpl{DB}
}

// FindAll implements repository.CustomerRepository.
func (e *CustomerRepoImpl) FindAll() (*[]models.Customer, error) {

	var customers []models.Customer
	err := e.DB.Joins("Nationality").
		Preload("FamilyLists").Find(&customers).Error
	if err != nil {
		fmt.Printf("[CutomerRepoImpl.ReadAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return &customers, nil
}

// ReadById implements repository.CustomerRepository.
func (e *CustomerRepoImpl) ReadById(id int64) (*models.Customer, error) {

	var customer models.Customer
	err := config.DB.Where("customers.customer_id = ?", id).
		Joins("Nationality").
		Preload("FamilyLists").
		First(&customer).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}

	return &customer, nil
}
