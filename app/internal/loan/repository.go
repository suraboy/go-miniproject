package loan

import (
	"gorm.io/gorm"

	"github.com/suraboy/go-miniproject/app/internal/models"
)

// Repository interface for loan operations
type Repository interface {
	Create(loan *models.Loan) error
	GetByID(id string) (*models.Loan, error)
	GetByEmail(email string) ([]*models.Loan, error)
	Update(loan *models.Loan) error
	Delete(id string) error
	GetAll(limit, offset int) ([]*models.Loan, error)
	GetByStatus(status string, limit, offset int) ([]*models.Loan, error)
}

// repository implements Repository interface
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new loan repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new loan application
func (r *repository) Create(loan *models.Loan) error {
	return r.db.Create(loan).Error
}

// GetByID retrieves a loan by ID
func (r *repository) GetByID(id string) (*models.Loan, error) {
	var loan models.Loan
	err := r.db.Where("id = ?", id).First(&loan).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

// GetByEmail retrieves loans by email
func (r *repository) GetByEmail(email string) ([]*models.Loan, error) {
	var loans []*models.Loan
	err := r.db.Where("email = ?", email).Find(&loans).Error
	return loans, err
}

// Update updates a loan application
func (r *repository) Update(loan *models.Loan) error {
	return r.db.Save(loan).Error
}

// Delete soft deletes a loan application
func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Loan{}, "id = ?", id).Error
}

// GetAll retrieves all loans with pagination
func (r *repository) GetAll(limit, offset int) ([]*models.Loan, error) {
	var loans []*models.Loan
	err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&loans).Error
	return loans, err
}

// GetByStatus retrieves loans by status with pagination
func (r *repository) GetByStatus(status string, limit, offset int) ([]*models.Loan, error) {
	var loans []*models.Loan
	err := r.db.Where("status = ?", status).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&loans).Error
	return loans, err
}
