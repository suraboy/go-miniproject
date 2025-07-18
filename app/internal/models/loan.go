package models

import (
	"time"

	"gorm.io/gorm"
)

// Loan represents the loan application in database
type Loan struct {
	ID            string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FullName      string         `gorm:"type:varchar(255);not null" json:"fullName"`
	MonthlyIncome float64        `gorm:"type:decimal(12,2);not null" json:"monthlyIncome"`
	LoanAmount    float64        `gorm:"type:decimal(12,2);not null" json:"loanAmount"`
	LoanPurpose   string         `gorm:"type:varchar(100);not null" json:"loanPurpose"`
	Age           int            `gorm:"not null" json:"age"`
	PhoneNumber   string         `gorm:"type:varchar(20);not null" json:"phoneNumber"`
	Email         string         `gorm:"type:varchar(255);not null" json:"email"`
	Status        string         `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	Message       string         `gorm:"type:text" json:"message"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Loan model
func (Loan) TableName() string {
	return "loans"
}

// LoanStatus constants
const (
	LoanStatusPending  = "pending"
	LoanStatusApproved = "approved"
	LoanStatusRejected = "rejected"
	LoanStatusReview   = "under_review"
)
