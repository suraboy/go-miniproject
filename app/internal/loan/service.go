package loan

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/suraboy/go-miniproject/app/internal/models"
)

type Service interface {
	ProcessLoanApplication(req LoanApplication) (*LoanApplicationResponse, error)
	GetLoanByID(id string) (*models.Loan, error)
	GetLoansByEmail(email string) ([]*models.Loan, error)
}

type service struct {
	repo Repository
}

func NewService() Service {
	return &service{
		// repo: repo,
	}
}

// ProcessLoanApplication processes a loan application and returns a response
func (s *service) ProcessLoanApplication(req LoanApplication) (*LoanApplicationResponse, error) {
	// Validate request
	if err := s.validateLoanApplication(req); err != nil {
		return nil, err
	}

	// Generate unique ID for the application
	applicationID := uuid.New().String()

	// Simple loan approval logic based on income and loan amount
	status := s.determineLoanStatus(req)
	message := s.getStatusMessage(status)

	// Create loan model for database
	loan := &models.Loan{
		ID:            applicationID,
		FullName:      req.FullName,
		MonthlyIncome: req.MonthlyIncome,
		LoanAmount:    req.LoanAmount,
		LoanPurpose:   req.LoanPurpose,
		Age:           req.Age,
		PhoneNumber:   req.PhoneNumber,
		Email:         req.Email,
		Status:        string(status),
		Message:       message,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Save to database
	// if err := s.repo.Create(loan); err != nil {
	// 	return nil, err
	// }

	response := &LoanApplicationResponse{
		ID:        applicationID,
		Status:    string(status),
		Message:   message,
		CreatedAt: loan.CreatedAt,
	}

	return response, nil
}

func (s *service) determineLoanStatus(req LoanApplication) LoanStatus {
	maxLoanAmount := req.MonthlyIncome * 5

	if req.LoanAmount <= maxLoanAmount && req.MonthlyIncome >= 3000 {
		return StatusApproved
	} else if req.MonthlyIncome >= 2000 {
		return StatusReview
	}

	return StatusRejected
}

// getStatusMessage returns appropriate message based on status
func (s *service) getStatusMessage(status LoanStatus) string {
	switch status {
	case StatusApproved:
		return "Congratulations! Your loan application has been approved."
	case StatusReview:
		return "Your loan application is under review. We will contact you within 3-5 business days."
	case StatusRejected:
		return "Unfortunately, your loan application has been rejected. Please contact us for more information."
	default:
		return "Your loan application has been received and is being processed."
	}
}

// validateLoanApplication validates the loan application data
func (s *service) validateLoanApplication(req LoanApplication) error {
	if req.FullName == "" {
		return errors.New("full name is required")
	}
	if req.MonthlyIncome <= 0 {
		return errors.New("monthly income must be greater than 0")
	}
	if req.LoanAmount <= 0 {
		return errors.New("loan amount must be greater than 0")
	}
	if req.Age < 18 {
		return errors.New("applicant must be at least 18 years old")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	if req.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	return nil
}

// GetLoanByID retrieves a loan by ID
func (s *service) GetLoanByID(id string) (*models.Loan, error) {
	return s.repo.GetByID(id)
}

// GetLoansByEmail retrieves loans by email
func (s *service) GetLoansByEmail(email string) ([]*models.Loan, error) {
	return s.repo.GetByEmail(email)
}
