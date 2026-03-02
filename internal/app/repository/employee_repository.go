package repository

import (
	"github.com/te-fa-bene/api-go/internal/app/domain"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) FindActiveByStoreAndEmail(storeID, email string) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.
		Where("store_id = ? AND email = ? AND deleted_at IS NULL AND is_active = true", storeID, email).
		First(&employee).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *EmployeeRepository) FindActiveByIDAndStore(employeeID, storeID string) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.
		Where("id = ? AND store_id = ? AND deleted_at IS NULL AND is_active = true", employeeID, storeID).
		First(&employee).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}
