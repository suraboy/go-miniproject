package loan

import "time"

// LoanApplication represents a loan application request
type LoanApplication struct {
	FullName      string  `json:"fullName" validate:"required"`
	MonthlyIncome float64 `json:"monthlyIncome" validate:"required,min=0"`
	LoanAmount    float64 `json:"loanAmount" validate:"required,min=0"`
	LoanPurpose   string  `json:"loanPurpose" validate:"required"`
	Age           int     `json:"age" validate:"required,min=18,max=100"`
	PhoneNumber   string  `json:"phoneNumber" validate:"required"`
	Email         string  `json:"email" validate:"required,email"`
}

// LoanApplicationResponse represents the response after submitting a loan application
type LoanApplicationResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// LoanStatus represents possible loan application statuses
type LoanStatus string

const (
	StatusPending  LoanStatus = "pending"
	StatusApproved LoanStatus = "approved"
	StatusRejected LoanStatus = "rejected"
	StatusReview   LoanStatus = "under_review"
)
