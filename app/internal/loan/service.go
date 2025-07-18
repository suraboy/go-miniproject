package loan

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	ProcessLoanApplication(req LoanApplication) (*LoanApplicationResponse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

// ProcessLoanApplication processes a loan application and returns a response
func (s *service) ProcessLoanApplication(req LoanApplication) (*LoanApplicationResponse, error) {
	// Generate unique ID for the application
	applicationID := uuid.New().String()

	// Simple loan approval logic based on income and loan amount
	status := s.determineLoanStatus(req)

	response := &LoanApplicationResponse{
		ID:        applicationID,
		Status:    string(status),
		Message:   s.getStatusMessage(status),
		CreatedAt: time.Now(),
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
